package main

import (
    "errors"
    "github.com/aws/aws-lambda-go/events"
    "github.com/stretchr/testify/assert"
    "github.com/vahrennd/ip-lookup/src/iplookup/model"
    "net/http"
    "testing"
)

func TestMain_handleInternal(t *testing.T) {
    request := events.APIGatewayProxyRequest{QueryStringParameters: make(map[string]string)}
    request.QueryStringParameters["address"] = "address"

    response, err := handleInternal(request, LookupMock{})

    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, response.StatusCode)
    assert.Contains(t, response.Body, "Results for address")
}

func TestMain_handleInternal_noAddress(t *testing.T) {
    request := events.APIGatewayProxyRequest{QueryStringParameters: make(map[string]string)}

    response, _ := handleInternal(request, LookupMock{})

    assert.Equal(t, http.StatusOK, response.StatusCode)
    assert.Contains(t, response.Body, "The address parameter is required.")
}

func TestMain_handleInternal_apiError(t *testing.T) {
    request := events.APIGatewayProxyRequest{QueryStringParameters: make(map[string]string)}
    request.QueryStringParameters["address"] = "whoops"

    _, err := handleInternal(request, LookupMock{})

    assert.Error(t, err, "The address parameter is required.")
}

type LookupMock struct{}

func (LookupMock) LookupDomain(address string) (model.LookupResponse, error) {
    if "whoops" == address {
        return model.LookupResponse{}, errors.New("whoops")
    }
    return model.LookupResponse{Whois: "whois"}, nil
}
