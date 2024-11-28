package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerSeven(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/ceaserCipher", handlers.CeaserCipherHandler)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}

