package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerSeven(w http.ResponseWriter, r *http.Request) {
	handlers.CeaserCipher(w, r)
}

