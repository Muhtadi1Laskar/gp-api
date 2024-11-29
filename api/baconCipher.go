package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerBaconCipher(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/baconCipher", handlers.BaconCipherHandler)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}