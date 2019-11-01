package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
)


func main() {
port := os.Getenv("PORT")
if port == "" {
        port = "8080"
        log.Printf("Defaulting to port %s", port)
}

host, _ := os.Hostname()
log.Printf("%s port %s vers. %s", host, port, version)

log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
