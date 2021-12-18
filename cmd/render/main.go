package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/hara/roomheatmap"
)

func main() {

	lambda.Start(handle)
}

func handle(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	data := request.PathParameters["data"]

	ptn := regexp.MustCompile(`^(-?\d{1,3}(?:\.\d)?)-(\d{1,3}(?:\.\d)?)$`)
	m := ptn.FindStringSubmatch(data)

	if m == nil {
		fmt.Printf("invalid data: %v\n", data)
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	t, _ := strconv.ParseFloat(m[1], 32)
	h, _ := strconv.ParseFloat(m[2], 32)

	heatmap := &roomheatmap.HeatMap{
		Room1: &roomheatmap.Metric{Temperature: float32(t), Humidity: float32(h)},
	}

	fmt.Printf("room1: %v\n", heatmap.Room1)

	result, err := heatmap.Render()
	if err != nil {
		fmt.Printf("could not render: %v\n", err)
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		IsBase64Encoded: true,
		StatusCode:      http.StatusOK,
		Body:            base64.StdEncoding.EncodeToString(result),
		Headers: map[string]string{
			"content-type": "image/png",
		},
	}, nil
}
