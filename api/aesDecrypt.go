package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerTwo(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/aesDecrypt", handlers.DecryptAES)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}
