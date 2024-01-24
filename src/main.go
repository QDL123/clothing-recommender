package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"fmt"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	// Destructure context and extract relevant data (ID of subscriber)
	return fmt.Sprintf("Hello World!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
