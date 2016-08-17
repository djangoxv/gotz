package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
)

func UtcIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "POSTing an iana db location will tell you current date in corresponding timezone")
    fmt.Fprintln(w, "Example: curl -X POST -d \"America/New_York\" http://" + r.Host +"/")
    fmt.Fprintln(w, TzConvert("America/New_York"))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World")
}

func TzTime(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    tz := TzConvert(string(body))
    fmt.Fprintln(w, tz)
}
