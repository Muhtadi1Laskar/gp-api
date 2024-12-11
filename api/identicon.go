package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerIdenticon(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/identicon", handlers.IdenticonHandler)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}