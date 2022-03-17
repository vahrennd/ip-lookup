package model

// represents a collection of all the various lookups we do - or plan to do
type LookupResponse struct {
    Whois string
    GeoIp GeoIp
}
