package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"go-stock-tracker/internal"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context) error {
	err := internal.NotifyClosePrice()
	if err != nil {
		return err
	}
	return nil
}
