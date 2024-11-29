package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerRot13(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/rot13Cipher", handlers.Rot13Handler)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}