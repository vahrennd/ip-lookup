package api

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "github.com/vahrennd/ip-lookup/src/iplookup/model"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestApi_testLookupGeoIp(t *testing.T) {
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "{\"Status\":\"success\",\"City\":\"KCMO\"}")
    }))
    lookupResponse := model.LookupResponse{}

    lookupGeoIp("1.2.3.4", &lookupResponse, mockServer.URL+"/")

    assert.Equal(t, "success", lookupResponse.GeoIp.Status)
    assert.Equal(t, "KCMO", lookupResponse.GeoIp.City)
}

func TestApi_testLookupGeoIpInvalidResponse(t *testing.T) {
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "you dialed the wrong number")
    }))
    lookupResponse := model.LookupResponse{}

    lookupGeoIp("1.2.3.4", &lookupResponse, mockServer.URL+"/")

    assert.NotEqual(t, "success", lookupResponse.GeoIp.Status)
}
