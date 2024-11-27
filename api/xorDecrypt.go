package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerSix(w http.ResponseWriter, r *http.Request) {
	handlers.XORDecrypt(w, r)
}

