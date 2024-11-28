package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerNine(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/vignereCipher", handlers.VigenereCipherHandler)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}