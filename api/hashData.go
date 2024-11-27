package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerOne(w http.ResponseWriter, r *http.Request) {
	handlers.HashData(w, r)
}
