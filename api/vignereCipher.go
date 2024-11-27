package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerEight(w http.ResponseWriter, r *http.Request) {
	handlers.VigenereCipherHandler(w, r)
}