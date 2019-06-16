package main

import (
	"context"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/aws/aws-lambda-go/lambda"
)

// Input takes in a secret
type Input struct {
	Secret string `json:"secret"`
}

// HandleRequest will be called when the lambda function is invoked
// it takes an input and checks if it matches our super secret value
func HandleRequest(ctx context.Context, input Input) (string, error) {
	log.SetHandler(text.New(os.Stderr))

	log.WithField("secret", input.Secret).Info("secret guessed")

	if input.Secret == "klaatu barada nikto" {
		return "secret guessed!", nil
	}
	return "try again", nil
}

func main() {
	lambda.Start(HandleRequest)
}
