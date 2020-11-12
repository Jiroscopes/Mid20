package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// EmailPass stored in env var
var EmailPass string

func send(body string) {
	from := "steven@midtwenty.com"
	pass := EmailPass
	to := "steven@midtwenty.com"

	log.Println(EmailPass)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:465",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent")
}

// Handler takes care of incoming requests
func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	send("Testing email!")

	name := request.QueryStringParameters["name"]
	response := fmt.Sprintf("Hello %s!", name)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "text/html; charset=UTF-8"},
		Body:       response,
	}, nil
}

// Init is ran before main()
func init() {
	//loads values from .env
	EmailPass, _ = os.LookupEnv("EMAILPASS")
}

func main() {
	// Initiate AWS Lambda handler
	lambda.Start(Handler)
}
