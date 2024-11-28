package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerFive(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/xorEncrypt", handlers.XOREncrypt)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}