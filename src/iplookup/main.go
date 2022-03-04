package main

import (
    "fmt"
    "github.com/vahrennd/ip-lookup/src/iplookup/api"
    "github.com/vahrennd/ip-lookup/src/iplookup/model"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/lookup", func(w http.ResponseWriter, r *http.Request) {
        addresses, ok := r.URL.Query()["address"]
        if !ok {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("The 'address' parameter is required."))
            return
        }

        address := addresses[0]

        Response, err := api.LookupIp(address)

        if err == nil {
            writeResponse(w, address, Response)
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "Failed to generate response for %q", address)
        }
    })

    log.Println("Listening on localhost:8080")

    log.Fatal(http.ListenAndServe(":8080", nil))
}

// writeResponse formats the results of any lookups performed and writes them back to the original requester.
func writeResponse(w http.ResponseWriter, address string, Response model.LookupResponse) {
    fmt.Fprintf(w, "Results for %q\n\n", address)
    fmt.Fprintf(w, "WHOIS:\n\n", address)
    w.Write([]byte(Response.Whois))
}
