package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerBase64(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/encodeBase64", handlers.BaseEncoderHandler)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}