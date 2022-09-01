package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Handler struct{
	token string
}

func NewHandler(token string) Handler {
	return Handler{ token: token }
}

func (h *Handler) Run(_ context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	filter := url.QueryEscape("stationLocation eq 'Clayton Hall'")
	reqUrl := fmt.Sprintf("https://api.tfgm.com/odata/Metrolinks?$filter=%s", filter)
	req, err := http.NewRequest("GET", reqUrl, nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", h.token)

	client := &http.Client{}

	res, err := client.Do(req)

	defer res.Body.Close()

	if err != nil {
		fmt.Print(err)
		return events.APIGatewayV2HTTPResponse{
			StatusCode: res.StatusCode,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       err.Error(),
		}, err
	}

	body, err := io.ReadAll(res.Body)

	return events.APIGatewayV2HTTPResponse{
		StatusCode: res.StatusCode,
		Headers:    map[string]string{"Content-Type": res.Header.Get("Content-Type")},
		Body: string(body),
	}, nil
}

func main() {
	token := os.Getenv("API_TOKEN")
	handler := NewHandler(token)

	lambda.Start(handler.Run)
}
