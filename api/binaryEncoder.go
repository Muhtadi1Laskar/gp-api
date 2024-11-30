package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerBinaryEncoder(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/binaryEncoder", handlers.BinaryEncoderHander)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}