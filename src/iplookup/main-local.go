package main

import (
    "fmt"
    "github.com/vahrennd/ip-lookup/src/iplookup/api"
    "github.com/vahrennd/ip-lookup/src/iplookup/utils"
    "log"
    "net/http"
)

// in case i forget to delete this before calling this done... i was using this file to run this app locally to make it
// easier to quickly test changes / debug. also because i'm not 100% sure that i'm not going to get a bill from AWS...
func main() {
    http.HandleFunc("/lookup", func(w http.ResponseWriter, r *http.Request) {
        addresses, ok := r.URL.Query()["address"]
        if !ok {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("The 'address' parameter is required."))
            return
        }

        address := addresses[0]

        response, err := api.LookupApi{}.LookupDomain(address)

        if err == nil {
            w.Write([]byte(utils.FormatResponse(address, response)))
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "Failed to generate response for %q", address)
        }
    })

    log.Println("Listening on localhost:8080")

    log.Fatal(http.ListenAndServe(":8080", nil))
}
