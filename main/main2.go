package main

import (
    "fmt"
    "log"
    "net/http"
    "reflect"
)

func handler(w http.ResponseWriter, r *http.Request) {
    text := "Hi there!!"
    fmt.Print(reflect.TypeOf(text))
    fmt.Fprintf(w, text)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":4444", nil))
}