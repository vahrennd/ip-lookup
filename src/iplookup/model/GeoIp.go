package model

import "fmt"

type GeoIp struct {
    Status      string
    Country     string
    CountryCode string
    Region      string
    RegionName  string
    City        string
    Zip         string
    Lat         float32
    Lon         float32
    Timezone    string
    Isp         string
    Org         string
    As          string
    Query       string
}

func (g GeoIp) String() string {
    return fmt.Sprintf(
        "Country=%s, CountryCode=%s, Region=%s, RegionName=%s, City=%s, Zip=%s, Lat=%f, Lon=%f, Timezone=%s, Isp=%s, Org=%s, As=%s, Query=%s",
        g.Country,
        g.CountryCode,
        g.Region,
        g.RegionName,
        g.City,
        g.Zip,
        g.Lat,
        g.Lon,
        g.Timezone,
        g.Isp,
        g.Org,
        g.As,
        g.Query,
    )
}
