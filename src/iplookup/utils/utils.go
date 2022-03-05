package utils

import (
    "fmt"
    "github.com/vahrennd/ip-lookup/src/iplookup/model"
)

const separator = "\n\n----------\n\n"

// FormatResponse formats the results of any lookups performed in a human-readable report
func FormatResponse(address string, Response model.LookupResponse) string {
    // TODO inefficient, maybe move to LookupResponse?
    var formattedResponse string

    formattedResponse += fmt.Sprintf("Results for %s", address)

    formattedResponse += separator

    formattedResponse += fmt.Sprintf("WHOIS:\n\n")
    formattedResponse += Response.Whois

    formattedResponse += separator

    if "success" == Response.GeoIp.Status {
        formattedResponse += fmt.Sprintf("GeoIP:\n\n")
        formattedResponse += fmt.Sprintf("This address is located in %s\n\n", Response.GeoIp.RegionName)
        formattedResponse += fmt.Sprintf("Full GeoIP data: %s", Response.GeoIp)
    } else {
        formattedResponse += "No GeoIP data available."
    }

    formattedResponse += separator

    return formattedResponse
}
