package api

import (
    "encoding/json"
    "github.com/vahrennd/ip-lookup/src/iplookup/model"
    "io/ioutil"
    "net/http"
)

const geoIpApiAddress = "http://ip-api.com/json/"

type LookupInterface interface {
    LookupDomain(address string) (model.LookupResponse, error)
}

type LookupApi struct{}

// LookupDomain performs various lookups about the given domain and returns a LookupResponse containing each one.
func (LookupApi) LookupDomain(address string) (model.LookupResponse, error) {
    return lookupDomainInternal(address, model.LookupWhois{}, geoIpApiAddress)
}

func lookupDomainInternal(address string, whois model.WhoisInterface, apiAddress string) (model.LookupResponse, error) {
    var lookupResponse model.LookupResponse
    var err error
    lookupResponse.Whois, err = whois.Whois(address)
    lookupGeoIp(address, &lookupResponse, apiAddress)
    return lookupResponse, err
}

func lookupGeoIp(address string, lookupResponse *model.LookupResponse, apiAddress string) {
    // looks like there are options out there if I wanted to pull in a dependency to do this, but I wanted to try it
    // the old-fashioned way. also, probably not necessary to pass lookupResponse by reference, but I also wanted to try that.
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
