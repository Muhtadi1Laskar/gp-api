package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
)

type AESRequestBody struct {
	Message string `json:"message"`
	Key string `json:"key"`
}

type AESDecryptRequest struct {
	CipherText string `json:"ciphertext"`
	Key string `json:"key"`
}

type AESResponseBody struct {
	CipherText string `json:"ciphertext"`
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
		CipherText: cipherText,
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