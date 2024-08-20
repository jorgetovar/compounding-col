package GetItemsFunction

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"os"
)

var (
	dynamoDbClient *dynamodb.DynamoDB
	tableName      = os.Getenv("COMPOUNDING_TABLE_NAME")
)

func init() {
	sess := session.Must(session.NewSession())
	dynamoDbClient = dynamodb.New(sess)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Example: Fetching all items from DynamoDB
	result, err := dynamoDbClient.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		log.Printf("Failed to scan items: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf(`{"error": "Failed to fetch items: %v"}`, err),
		}, nil
	}

	// Convert result to JSON
	jsonResponse, err := json.Marshal(result.Items)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       `{"error": "Failed to parse items"}`,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonResponse),
	}, nil
}

func main() {
	lambda.Start(handler)
}
