package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGetRecommendation(t *testing.T) {
	err := GetRecommendation(os.Getenv("LOCATION"), os.Getenv("PHONE_NUMBER"))
	if err != nil {
		t.Errorf("Failed to get temperature: %v", err)
	}
	// You can add more checks here based on your expectations
	fmt.Printf("DONE!")
}
