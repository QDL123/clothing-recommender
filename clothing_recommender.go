package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type WeatherResponse struct {
	Timelines struct {
		Daily []struct {
			Values struct {
				TemperatureApparentAvg float64
				TemperatureApparentMin float64
				TemperatureApparentMax float64
				RainIntensityMax       float64
			}
		}
	}
}

func getWeatherData(location string) (WeatherResponse, error) {
	token := os.Getenv("TOKEN")
	baseURL, err := url.Parse(os.Getenv("WEATHER_URL"))

	// URL values
	params := url.Values{}
	params.Add("location", location)
	params.Add("apikey", token)
	params.Add("timesteps", "daily")
	params.Add("units", "imperial")
	baseURL.RawQuery = params.Encode()
	resp, err := http.Get(baseURL.String())

	var weatherResponse WeatherResponse
	if err != nil {
		return weatherResponse, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return weatherResponse, err
	}

	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return weatherResponse, err
	}
	return weatherResponse, nil
}

func generateRecommendationText(data WeatherResponse) (text string) {
	avgFeelsLike := data.Timelines.Daily[0].Values.TemperatureApparentAvg
	maxRainIntensity := data.Timelines.Daily[0].Values.RainIntensityMax

	recommendation := ""
	if avgFeelsLike > 70 {
		recommendation += "Today will be hot! I recommend wearing shorts and a T-shirt."
	} else if avgFeelsLike > 60 {
		recommendation += "Today will be warm!. I recommend pants and a T-shirt."
	} else if avgFeelsLike > 50 {
		recommendation += "Today will be relatively cool. I recommend wearing pants and a T-shirt plus a light layer like a quarter zip or light sweater."
	} else if avgFeelsLike > 40 {
		recommendation += "I recommend wearing a light jacket and pants."
	} else {
		recommendation += "Today will be cold. Where pants and a thick jacket or wear multiple layers."
	}

	if maxRainIntensity > 0.4 {
		recommendation += " Today will be rainy. Wear a rain jacket!"
	} else {
		recommendation += " No need to worry about precipitation today."
	}

	return recommendation
}

func sendRecommendation(text string, phoneNumber string) error {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody(text)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
		return err
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

	return nil
}

func GetRecommendation(location string, phoneNumber string) error {
	// retrieve weather data
	weatherData, err := getWeatherData(location)
	if err != nil {
		log.Printf("An error occurred: %s\n", err)
		return err
	}

	// generate recommendation string
	text := generateRecommendationText(weatherData)

	// send recommendation via text
	err = sendRecommendation(text, phoneNumber)
	if err != nil {
		log.Printf("An error occurred: %s\n", err)
	}
	fmt.Println(text)
	return nil
}
