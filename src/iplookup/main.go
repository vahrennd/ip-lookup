package main

import (
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

func handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return handleInternal(request, api.LookupApi{})
}

func handleInternal(request events.APIGatewayProxyRequest, lookupApi api.LookupInterface) (events.APIGatewayProxyResponse, error) {
    address := request.QueryStringParameters["address"]
    if address == "" {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusOK,
            Body:       "The address parameter is required.",
        }, nil
    }

    response, err := lookupApi.LookupDomain(address)

    if err == nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusOK,
            Body:       utils.FormatResponse(address, response),
        }, nil
    } else {
        return events.APIGatewayProxyResponse{}, errors.New("failed to generate response")
    }
}
