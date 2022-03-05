package api

import (
    "errors"
    "fmt"
    "github.com/stretchr/testify/assert"
    "github.com/vahrennd/ip-lookup/src/iplookup/model"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestLookupApi_LookupDomain(t *testing.T) {
    mockGeoIpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "{\"Status\":\"success\",\"City\":\"KCMO\"}")
    }))

    lookupResponse, err := lookupDomainInternal("address", MockWhois{}, mockGeoIpServer.URL+"/")

    assert.NoError(t, err)
    assert.Equal(t, "whois", lookupResponse.Whois)
    assert.Equal(t, "success", lookupResponse.GeoIp.Status)
    assert.Equal(t, "KCMO", lookupResponse.GeoIp.City)
}

func TestLookupApi_LookupDomain_whoisError(t *testing.T) {
    mockGeoIpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "{\"Status\":\"success\",\"City\":\"KCMO\"}")
    }))

    lookupResponse, err := lookupDomainInternal("whoops", MockWhois{}, mockGeoIpServer.URL+"/")

    assert.Error(t, err)
    assert.Equal(t, "", lookupResponse.Whois)
    assert.Equal(t, "success", lookupResponse.GeoIp.Status)
    assert.Equal(t, "KCMO", lookupResponse.GeoIp.City)
}

func TestApi_testLookupGeoIp_invalidResponse(t *testing.T) {
    mockGeoIpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "you dialed the wrong number")
    }))
    lookupResponse := model.LookupResponse{}

    lookupGeoIp("1.2.3.4", &lookupResponse, mockGeoIpServer.URL+"/")

    assert.NotEqual(t, "success", lookupResponse.GeoIp.Status)
}

type MockWhois struct{}

func (MockWhois) Whois(domain string) (string, error) {
    if "whoops" == domain {
        return "", errors.New("whoops")
    }
    return "whois", nil
}
