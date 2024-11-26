package main

import (
	"go-api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/aes/encrypt", handlers.EncryptAES)
	http.ListenAndServe(":5000", nil)
}