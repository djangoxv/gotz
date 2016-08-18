// gotz/main.go
package main

import (
    "flag"
    "log"
    "net/http"
    "os"
)

// creates a logger 
var (
    logError *log.Logger
)

// the -logfile flag creates a log file
func Init(logfile string) {
    
    // open a log file, if you can
    file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Fatal Error: ", err)
    }
    
    logError = log.New(file, "INFO: ", log.Ldate|log.Ltime)
}


func main() {
    // process flags passed from service startup
    port      := flag.String("port", "80", "Port to bind gopkg to")
    logfile   := flag.String("logfile", "/tmp/gotz.log", "Log file path and name")
    flag.Parse()

    // begin logging
    Init(*logfile)

    router := NewRouter()
    logError.Fatal(http.ListenAndServe(":" + *port, router))
}
