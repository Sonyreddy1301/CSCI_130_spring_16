package main

import (
    "io"
    "net/http"
    "log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, world!\n")
}

func main() {
    http.HandleFunc("/hello", HelloServer)
    err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

//Windows User use commands 

//For generating private Key
//openssl genrsa -out server.key 2048

//To generate certificate
//openssl req -new -x509 -key server.key -out server.pem -days 3650
