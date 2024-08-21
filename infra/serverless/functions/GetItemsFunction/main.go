package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

// Response is the structure for the response JSON
type Response struct {
	Message string `json:"message"`
}

func response(code int, object interface{}) events.APIGatewayV2HTTPResponse {
	marshalled, err := json.Marshal(object)
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error())
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: code,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:            string(marshalled),
		IsBase64Encoded: false,
	}
}

func errResponse(status int, body string) events.APIGatewayV2HTTPResponse {
	message := map[string]string{
		"message": body,
	}

	messageBytes, _ := json.Marshal(&message)
	fmt.Println("Error: ", string(messageBytes))
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(messageBytes),
	}

}

// HelloHandler handles API Gateway requests and returns "Hello, World!"
func HelloHandler(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	fmt.Println("HelloHandler invoked")
	helloMessage := Response{
		Message: "Hello, World!",
	}
	httpResponse := response(http.StatusOK, helloMessage)
	fmt.Println("Response: ", httpResponse)
	return httpResponse, nil
}

func main() {
	fmt.Println("Starting GetItemsFunction")
	lambda.Start(HelloHandler)
}
