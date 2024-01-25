package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetRecommendation(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		fmt.Print("No env file found")
		return
	}
	err := GetRecommendation(os.Getenv("LOCATION"), os.Getenv("PHONE_NUMBER"))
	if err != nil {
		t.Errorf("Failed to get temperature: %v", err)
	}
	// You can add more checks here based on your expectations
	fmt.Printf("DONE!")
}
