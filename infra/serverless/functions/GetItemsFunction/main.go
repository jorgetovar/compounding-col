package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is the structure for the response JSON
type Response struct {
	Message string `json:"message"`
}

func HelloHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Create a response object
	helloMessage := Response{
		Message: "Hello, World!",
	}

	// Marshal the response object to JSON
	body, err := json.Marshal(helloMessage)
	if err != nil {
		// Return a 500 Internal Server Error response if JSON marshaling fails
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to marshal response",
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	// Return the correct API Gateway proxy response
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	lambda.Start(HelloHandler)
}
