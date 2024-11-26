package handlers

import (
	"go-api/Hashes"
	"net/http"
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

	hashedData, err := hashes.HashData(requestBody.Message, requestBody.Hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := HashResponseBody{
		Hash: hashedData,
	}

	writeJSONResponse(w, http.StatusOK, response)
}
