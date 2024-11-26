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

type VerifyHashRequest struct {
	OldHash string `json:"oldhash"`
	Message string `json:"message"`
	Hash string `json:"hash"`
}

type VerifyHashResponse struct {
	Altered bool `json:"altered"`
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

func VerifyData(w http.ResponseWriter, r *http.Request) {
	var requestBody VerifyHashRequest

	if err := readRequestBody(r, &requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	isAltered, err := hashes.VerifyHash(requestBody.OldHash, requestBody.Message, requestBody.Hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody := VerifyHashResponse {
		Altered: isAltered,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}
