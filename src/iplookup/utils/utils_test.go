package utils

import (
    "github.com/stretchr/testify/assert"
    "github.com/vahrennd/ip-lookup/src/iplookup/model"
    "testing"
)

func TestUtils_testFormatResponse(t *testing.T) {
    geoIp := model.GeoIp{Status: "success", City: "KCMO", RegionName: "Missouri"}
    lookupResponse := model.LookupResponse{Whois: "definitely.real.whois.data", GeoIp: geoIp}

    report := FormatResponse("address", lookupResponse)

    // not testing for an exact match (would be too rigid if small parts of the report get reformatted), so we'll
    // just assert that the important parts show up appropriately
    assert.Contains(t, report, "Results for address")
    assert.Contains(t, report, "definitely.real.whois.data")
    assert.Contains(t, report, "This address is located in Missouri")
    assert.Contains(t, report, "Full GeoIP data:")
    assert.Contains(t, report, "City=KCMO")
}

func TestUtils_testFormatResponse_noGeoIpData(t *testing.T) {
    geoIp := model.GeoIp{Status: "not-success"}
    lookupResponse := model.LookupResponse{Whois: "definitely.real.whois.data", GeoIp: geoIp}

    report := FormatResponse("address", lookupResponse)

    assert.Contains(t, report, "No GeoIP data available.")

    // make sure none of the GeoIp report is printed if we didn't successfully look that up
    assert.NotContains(t, report, "This address is located in")
    assert.NotContains(t, report, "Full GeoIP data:")
    assert.NotContains(t, report, "City=")
}
