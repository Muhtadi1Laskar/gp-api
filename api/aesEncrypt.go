package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerThree(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/aesEncrypt", handlers.CipherAES)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}
