package model

import "github.com/likexian/whois"

type WhoisInterface interface {
    Whois(domain string) (string, error)
}

type LookupWhois struct{}

func (LookupWhois) Whois(domain string) (string, error) {
    return whois.Whois(domain)
}
