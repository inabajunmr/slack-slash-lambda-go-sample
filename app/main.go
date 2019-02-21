package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// valification token
	token := os.Getenv("VERIFICATION_TOKEN")
	vals, _ := url.ParseQuery(request.Body)
	if vals.Get("token") != token {
		// invalid token
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Invalid token."),
			StatusCode: 200,
		}, nil
	}

	// ### Write your command logic ###

	return events.APIGatewayProxyResponse{
		Body:       "Hello world.",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
