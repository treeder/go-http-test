package main

import (
    "flag"
    "http"
    "io"
    "log"
    "strconv"
)

var env = flag.String("e", "development", "environment")

func main() {
    flag.Parse()

    // Start up our http server
    httpPort := 8080
    if *env != "development" {
        httpPort = 80
    }

    http.Handle("/", http.HandlerFunc(Ping))
    err := http.ListenAndServe(":"+strconv.Itoa(httpPort), nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func Ping(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "pong\n")
}
