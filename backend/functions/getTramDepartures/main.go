package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Handler struct {

}

func (h *Handler) Run(_ context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	handler := Handler{}

	lambda.Start(handler.Run)
}