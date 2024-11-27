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

type VerifyHashRequest struct {
	OldHash string `json:"oldhash"`
	Message string `json:"message"`
	Hash string `json:"hash"`
}

type VerifyHashResponse struct {
	IsSame bool `json:"issame"`
}

func HashData(w http.ResponseWriter, r *http.Request) {
	var requestBody HashRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	hashedData, err := hashes.HashData(requestBody.Message, requestBody.Hash)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
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
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	changeStatus, err := hashes.VerifyHash(requestBody.OldHash, requestBody.Message, requestBody.Hash)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseBody := VerifyHashResponse {
		IsSame: changeStatus,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}
