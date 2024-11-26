package handlers

import (
	"go-api/Hashes"
	"net/http"
	"strings"
)

type HashRequestBody struct {
	Message string `json:"message"`
	Hash    string `json:"hash"`
}

type HashResponseBody struct {
	Hash string `json:"hash"`
}

func HashData(w http.ResponseWriter, r *http.Request) {
	var requestBody HashRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	formattedStr := strings.ToLower(requestBody.Hash)
	hashedData, err := hashes.HashData(requestBody.Message, formattedStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := HashResponseBody{
		Hash: hashedData,
	}

	writeJSONResponse(w, http.StatusOK, response)
}
