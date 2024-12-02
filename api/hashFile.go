package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerHashFile(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/hashFile", handlers.HandleFileHash)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}