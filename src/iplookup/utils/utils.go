package utils

import (
    "fmt"
    "github.com/vahrennd/ip-lookup/src/iplookup/model"
    "strings"
)

const separator = "\n\n----------\n\n"

// FormatResponse formats the results of any lookups performed into a human-readable report
func FormatResponse(address string, response model.LookupResponse) string {
    var sb strings.Builder

    sb.WriteString(fmt.Sprintf("Results for %s", address))

    sb.WriteString(separator)

    if response.Whois != "" {
        sb.WriteString(fmt.Sprintf("WHOIS:\n\n"))
        sb.WriteString(response.Whois)
    } else {
        sb.WriteString("No Whois data available.")
    }

    sb.WriteString(separator)

    if "success" == response.GeoIp.Status {
        sb.WriteString(fmt.Sprintf("GeoIP:\n\n"))
        sb.WriteString(fmt.Sprintf("This address is located in %s\n\n", response.GeoIp.RegionName))
        sb.WriteString(fmt.Sprintf("Full GeoIP data: %s", response.GeoIp))
    } else {
        sb.WriteString("No GeoIP data available.")
    }

    sb.WriteString(separator)

    return sb.String()
}
