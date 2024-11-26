package handler

import (
	"go-api/handlers"
	"net/http"
)

func Main() {
	http.HandleFunc("/", handlers.HashData)
	http.ListenAndServe(":8080", nil)
}
