package main

import (
    "fmt"
    "strings"
    "net/http"
    "github.com/pborman/uuid"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        uuidWithHyphen := uuid.NewRandom()
        uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
        fmt.Fprintf(w, "Welcome to my website!\n")
        fmt.Fprintf(w, uuid)
    })

    fmt.Print("Starting server...\n")

    http.ListenAndServe(":8080", nil)
}
