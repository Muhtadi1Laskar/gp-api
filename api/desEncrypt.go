package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerFour(w http.ResponseWriter, r *http.Request) {
	handlers.DecryptAES(w, r)
}