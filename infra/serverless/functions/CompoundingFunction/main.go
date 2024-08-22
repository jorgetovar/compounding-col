package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is the structure for the response JSON
type Response struct {
	Message      string    `json:"message"`
	GainsPerYear []float64 `json:"gainsPerYear"`
}

type Request struct {
	Principal  float64 `json:"principal"`
	AnnualRate float64 `json:"annualRate"`
	Years      int     `json:"years"`
}

func HelloHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var req Request
	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return createResponse(400, "Invalid request body")
	}
	fmt.Println("Request", req)
	gainsPerYear := CalculateCompoundInterest(req.Principal, req.AnnualRate, req.Years)
	fmt.Println(gainsPerYear)
	response := Response{
		Message:      "Calculation successful",
		GainsPerYear: gainsPerYear,
	}
	// Marshal the response object to JSON
	body, err := json.Marshal(response)
	if err != nil {
		return createResponse(500, "Error marshalling response")
	}

	// Return the correct API Gateway proxy response
	return createResponse(200, string(body))
}

func createResponse(statusCode int, body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	lambda.Start(HelloHandler)
}
