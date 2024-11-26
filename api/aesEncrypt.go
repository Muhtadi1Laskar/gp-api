package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerThree(w http.ResponseWriter, r *http.Request) {
	handlers.EncryptAES(w, r)
}