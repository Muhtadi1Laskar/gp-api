package handler

import (
	"go-api/handlers"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handlers.HashData(w, r)
}

func Main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
