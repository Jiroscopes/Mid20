package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func makeResponse(statusCode int, respBody string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "text/html; charset=UTF-8"},
		Body:       respBody,
	}
}

func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	statusCode := 200
	response := "Hi Steven!"
	return makeResponse(statusCode, response), nil
}

func main() {
	// Initiate AWS Lambda handler
	lambda.Start(Handler)
}
