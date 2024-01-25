package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func HandleRequest(ctx context.Context, event interface{}) (string, error) {
	fmt.Println("event", event)
	fmt.Println("IN HANDLE REQUEST FUNCTION")
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
	log.Print("ENTERED MAIN FUNCTION")
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	lambda.Start(HandleRequest)
}
