package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// Message is the input to the function and
// includes a FirstName and LastName
type Message struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Response is sent back and contains a greeting
// string
type Response struct {
	Greeting string `json:"greeting"`
}

// HandleRequest will be called when the lambda function is invoked
// it takes a Message and returns a Response that contains a greeting
// this greeting contains the first and last name specified
func HandleRequest(ctx context.Context, m Message) (Response, error) {
	return Response{Greeting: fmt.Sprintf("Hello, %s %s", m.FirstName, m.LastName)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
