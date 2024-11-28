package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerFive(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/xorCipher", handlers.CipherXOR)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}