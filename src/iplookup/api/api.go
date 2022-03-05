package api

import (
    "encoding/json"
    "github.com/likexian/whois"
    "github.com/vahrennd/ip-lookup/src/iplookup/model"
    "io/ioutil"
    "net/http"
)

const geoIpApiAddress = "http://ip-api.com/json/"

// LookupIp performs various lookups about the given address and returns a LookupResponse containing each one.
func LookupIp(address string) (model.LookupResponse, error) {
    var lookupResponse model.LookupResponse
    var err error
    lookupResponse.Whois, err = whois.Whois(address)
    lookupGeoIp(address, &lookupResponse, geoIpApiAddress)
    return lookupResponse, err
}

func lookupGeoIp(address string, lookupResponse *model.LookupResponse, apiAddress string) {
    resp, err := http.Get(apiAddress + address)
    if err != nil {
        return
    }

    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return
    }

    err = json.Unmarshal(bytes, &lookupResponse.GeoIp)
    if err != nil {
        return
    }
}
