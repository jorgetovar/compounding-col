package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, request map[string]interface{}) (map[string]interface{}, error) {
	// Create a simple response message
	resp := Response{
		Message: "Hello, World!",
	}

	// Marshal the response into JSON
	body, err := json.Marshal(resp)
	if err != nil {
		return map[string]interface{}{
			"statusCode": http.StatusInternalServerError,
			"body":       "Internal Server Error",
		}, nil
	}

	// Return the response with status code 200
	return map[string]interface{}{
		"statusCode": http.StatusOK,
		"body":       string(body),
	}, nil
}

func main() {
	// Manually invoke the handler function since we're not using aws-lambda-go
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Convert the request to a map
		request := map[string]interface{}{}
		_, err := handler(context.Background(), request)
		if err != nil {
			return
		}
	})
	fmt.Println("Lambda function initialized and ready.")
}
