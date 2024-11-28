package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerOne(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/hashData", handlers.HashData)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
	// handlers.HashData(w, r)
}
