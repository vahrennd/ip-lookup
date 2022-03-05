package main

import (
    "context"
    "errors"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/vahrennd/ip-lookup/src/iplookup/api"
    "github.com/vahrennd/ip-lookup/src/iplookup/utils"
    "net/http"
)

func main() {
    lambda.Start(handle)
}

func handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    address := request.QueryStringParameters["address"]
    if address == "" {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusOK,
            Body:       "The address parameter is required.",
        }, nil
    }

    Response, err := api.LookupIp(address)

    if err == nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusOK,
            Body:       utils.FormatResponse(address, Response),
        }, nil
    } else {
        // TODO log error?
        return events.APIGatewayProxyResponse{}, errors.New("failed to generate response")
    }
}
