# ip-lookup

A simple tool that, when given a domain URL or IP address, will look up various information about it.

Current reports:
* Whois
* GeoIP

To build, run `make`.

To deploy an AWS lambda function with an API gateway, configure the AWS cli tool with your user and secret IDs then run `terraform apply`.

To use, call the URL returned from terraform then add the `?address=[your address]` parameter.