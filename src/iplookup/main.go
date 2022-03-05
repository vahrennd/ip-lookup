package main

import (
    "context"
    "errors"
    "fmt"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/vahrennd/ip-lookup/src/iplookup/api"
    "github.com/vahrennd/ip-lookup/src/iplookup/model"
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
            Body:       formatResponse(address, Response),
        }, nil
    } else {
        // TODO log error?
        return events.APIGatewayProxyResponse{}, errors.New("failed to generate response")
    }
}

// formatResponse formats the results of any lookups performed in a human-readable report
func formatResponse(address string, Response model.LookupResponse) string {
    // TODO inefficient, maybe move to LookupResponse?
    var formattedResponse string
    formattedResponse += fmt.Sprintf("Results for %q\n\n", address)
    formattedResponse += fmt.Sprintf("WHOIS:\n\n")
    formattedResponse += Response.Whois
    return formattedResponse
}
