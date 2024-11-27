package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
)

type XORRequestBody struct {
	Message string `json:"message"`
	Key     string `json:"key"`
}

type XORResponseBody struct {
	Message string `json:"message"`
}

func XOREncrypt(w http.ResponseWriter, r *http.Request) {
	var requestBody XORRequestBody
	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	cipherText := ciphers.EncryptXOR(requestBody.Message, requestBody.Key)

	responseBody := XORResponseBody{
		Message: cipherText,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}

func XORDecrypt(w http.ResponseWriter, r *http.Request) {
	var requestBody XORRequestBody
	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	plainText, err := ciphers.DecryptXOR(requestBody.Message, requestBody.Key)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseBody := XORResponseBody{
		Message: plainText,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}
