package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerAtbashCipher(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/atbashCipher", handlers.AtbashHandler)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}