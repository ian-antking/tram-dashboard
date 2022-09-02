package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	metrolink "github.com/ian-antking/tram-dashboard/backend/getTramDepartures/repository"
)

type Handler struct {
	token string
}

func NewHandler(token string) Handler {
	return Handler{token: token}
}

func (h *Handler) Run(_ context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	tramStop := event.PathParameters["tramStop"]
	filter := url.QueryEscape(fmt.Sprintf("stationLocation eq '%s'", tramStop))
	reqUrl := fmt.Sprintf("https://api.tfgm.com/odata/Metrolinks?$filter=%s", filter)
	req, err := http.NewRequest("GET", reqUrl, nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", h.token)

	client := &http.Client{}

	res, err := client.Do(req)

	defer res.Body.Close()

	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: res.StatusCode,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       err.Error(),
		}, err
	}

	body, err := io.ReadAll(res.Body)

	var metrolinkResponseBody metrolink.ResponseBody
	jsonErr := json.Unmarshal(body, &metrolinkResponseBody)

	if jsonErr != nil {
		fmt.Print(jsonErr.Error())
	}

	var trams = make([]metrolink.Tram, 0, len(metrolinkResponseBody.Value))

	for _, tram := range metrolinkResponseBody.Value {
		if tram.Dest0 != "" {
			trams = append(trams, metrolink.Tram{
				Destination: tram.Dest0,
				Carriages:   tram.Carriages0,
				Status:      tram.Status0,
				Wait:        tram.Wait0,
				LastUpdated: tram.LastUpdated,
				Message:     tram.MessageBoard,
			})
		}
		if tram.Dest1 != "" {
			trams = append(trams, metrolink.Tram{
				Destination: tram.Dest1,
				Carriages:   tram.Carriages1,
				Status:      tram.Status1,
				Wait:        tram.Wait1,
				LastUpdated: tram.LastUpdated,
				Message:     tram.MessageBoard,
			})
		}
		if tram.Dest2 != "" {
			trams = append(trams, metrolink.Tram{
				Destination: tram.Dest2,
				Carriages:   tram.Carriages2,
				Status:      tram.Status2,
				Wait:        tram.Wait2,
				LastUpdated: tram.LastUpdated,
				Message:     tram.MessageBoard,
			})
		}
		if tram.Dest3 != "" {
			trams = append(trams, metrolink.Tram{
				Destination: tram.Dest3,
				Carriages:   tram.Carriages3,
				Status:      tram.Status3,
				Wait:        tram.Wait3,
				LastUpdated: tram.LastUpdated,
				Message:     tram.MessageBoard,
			})
		}
	}

	responseBody, _ := json.Marshal(metrolink.TramsResponseBody{
		Trams: trams,
	})

	return events.APIGatewayV2HTTPResponse{
		StatusCode: res.StatusCode,
		Headers:    map[string]string{"Content-Type": res.Header.Get("Content-Type")},
		Body:       string(responseBody),
	}, nil
}

func main() {
	token := os.Getenv("API_TOKEN")
	handler := NewHandler(token)

	lambda.Start(handler.Run)
}
