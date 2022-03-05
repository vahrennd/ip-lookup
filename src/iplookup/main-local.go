package main

import (
    "fmt"
    "github.com/vahrennd/ip-lookup/src/iplookup/api"
    "github.com/vahrennd/ip-lookup/src/iplookup/utils"
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
            w.Write([]byte(utils.FormatResponse(address, Response)))
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "Failed to generate response for %q", address)
        }
    })

    log.Println("Listening on localhost:8080")

    log.Fatal(http.ListenAndServe(":8080", nil))
}
