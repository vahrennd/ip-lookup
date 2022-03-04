package api

import (
    "github.com/likexian/whois"
    "github.com/vahrennd/ip-lookup/src/iplookup/model"
)

// LookupIp performs various lookups about the given address and returns a LookupResponse containing each one.
func LookupIp(address string) (model.LookupResponse, error) {
    var Response model.LookupResponse
    var err error
    Response.Whois, err = whois.Whois(address)
    return Response, err
}
