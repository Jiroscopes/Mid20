package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// EmailPass stored in env var
var EmailPass string

// User represents the info filled in by site user on contact form
type User struct {
	FirstName    string `json:"FirstName"`
	BusinessName string `json:"BusinessName"`
	Email        string `json:"Email"`
	Comments     string `json:"Comments"`
}

func makeResponse(statusCode int, respBody string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "text/html; charset=UTF-8"},
		Body:       respBody,
	}
}

// Send the email
func send(body *User) error {
	from := "steven@midtwenty.com"
	pass := EmailPass
	to := "steven@midtwenty.com"

	// Message format must follow RFC 822-style https://golang.org/pkg/net/smtp/#SendMail
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Mid-Twenty Form Submission\n\n" +
		body.FirstName + "\n" + body.BusinessName + "\n" + body.Email + "\n" + body.Comments

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		return errors.New("Sendmail has failed")
	}

	return nil
}

// Handler takes care of incoming requests
func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	// Default to failure
	var response string = "Failed"
	var statusCode int = 400

	// If request is GET
	if strings.ToLower(request.HTTPMethod) == "get" {
		if request.Body == "" {
			return makeResponse(statusCode, response), nil
		}
	}

	// If request body is empty
	if request.Body == "" {
		return makeResponse(statusCode, response), nil
	}

	// Byte slice from string
	bytes := []byte(request.Body)

	var user User

	// Read []byte to struct
	err := json.Unmarshal(bytes, &user)

	if err != nil {
		return makeResponse(statusCode, response), nil
	}

	if user.FirstName == "" || user.BusinessName == "" || user.Email == "" {
		return makeResponse(statusCode, response), nil
	}

	err = send(&user)

	if err == nil {
		response = fmt.Sprintf("Success")
		statusCode = 200
	}

	return makeResponse(statusCode, response), nil
}

// Init is ran before main()
func init() {
	EmailPass, _ = os.LookupEnv("EMAILPASS")
}

func main() {
	// Initiate AWS Lambda handler
	lambda.Start(Handler)
}
