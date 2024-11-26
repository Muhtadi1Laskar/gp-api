package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
)

type AESRequestBody struct {
	Message string `json:"message"`
	Key string `json:"key"`
}

type AESResponseBody struct {
	CipherText string `json:"ciphertext"`
}

func EncryptAES(w http.ResponseWriter, r *http.Request) {
	var requestBody AESRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cipherText, err := ciphers.AESEncrypt(requestBody.Message, requestBody.Key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody := AESResponseBody{
		CipherText: cipherText,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}