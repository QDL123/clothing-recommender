package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	log.Print("Clothing-recommender lambda triggered\n")

	location := os.Getenv("LOCATION")
	phoneNumber := os.Getenv("PHONE_NUMBER")
	err := GetRecommendation(location, phoneNumber)

	if err != nil {
		return "An error occurred", err
	}

	return "Done", nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	lambda.Start(HandleRequest)
}
