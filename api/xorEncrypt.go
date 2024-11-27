package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerFive(w http.ResponseWriter, r *http.Request) {
	handlers.XOREncrypt(w, r)
}

