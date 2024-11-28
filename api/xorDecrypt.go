package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerSix(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/xorDecrypt", handlers.XORDecrypt)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}
