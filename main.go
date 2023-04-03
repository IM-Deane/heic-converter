package main

import (
	"log"
	"net/http"

	"github.com/IM-Deane/heic-converter/handlers"
)


func main() {
    http.HandleFunc("/", handlers.Health)
    http.HandleFunc("/api/convert", handlers.ConvertImage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}