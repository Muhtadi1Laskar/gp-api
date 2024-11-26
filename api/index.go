package handler

import (
	"go-api/handlers"
	"net/http"
)

func HashHandler(w http.ResponseWriter, r *http.Request) {
	handlers.HashData(w, r)
}

func Main() {
	http.HandleFunc("/", HashHandler)
	http.ListenAndServe(":8080", nil)
}
