package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerCipherFile(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/cipherFile", handlers.HandleCipherFile)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}