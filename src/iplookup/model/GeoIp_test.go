package model

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestGeoIp_String(t *testing.T) {
    geoIp := GeoIp{
        Status:      "a",
        Country:     "b",
        CountryCode: "c",
        Region:      "d",
        RegionName:  "e",
        City:        "f",
        Zip:         "g",
        Lat:         1,
        Lon:         2,
        Timezone:    "h",
        Isp:         "i",
        Org:         "j",
        As:          "k",
        Query:       "l",
    }
    expected := "Country=b, CountryCode=c, Region=d, RegionName=e, City=f, Zip=g, Lat=1.000000, Lon=2.000000, Timezone=h, Isp=i, Org=j, As=k, Query=l"
    assert.Equal(t, expected, fmt.Sprintf("%s", geoIp))
}
