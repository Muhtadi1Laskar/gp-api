package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
)

type AESRequestBody struct {
	Message string `json:"message" validate:"required"`
	Key string `json:"key" validate:"required"`
}

type AESDecryptRequest struct {
	CipherText string `json:"ciphertext" validate:"required"`
	Key string `json:"key" validate:"required"`
}

type AESResponseBody struct {
	Message string `json:"message"`
}

type AESDecryptResponse struct {
	Message string `json:"message"`
}

func EncryptAES(w http.ResponseWriter, r *http.Request) {
	var requestBody AESRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	cipherText, err := ciphers.AESEncrypt(requestBody.Message, requestBody.Key)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseBody := AESResponseBody{
		Message: cipherText,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}

func DecryptAES(w http.ResponseWriter, r *http.Request) {
	var requestBody AESDecryptRequest

	if err := readRequestBody(r, &requestBody); err != nil {
		http.Error(w, error.Error(err), http.StatusInternalServerError)
		return
	}

	plaintText, err := ciphers.AESDecrypt(requestBody.CipherText, requestBody.Key)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseBody := AESDecryptResponse{
		Message: plaintText,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}