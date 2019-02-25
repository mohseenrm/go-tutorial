package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hola from server!")
}

func main() {
    http.HandleFunc("/", rootHandler)
    http.ListenAndServe(":8080", nil)
}