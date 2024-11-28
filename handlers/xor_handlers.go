package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
	"strings"
)

type XORRequestBody struct {
	Message string `json:"message" validate:"required"`
	Key     string `json:"key" validate:"required"`
	Type     string `json:"type" validate:"required"`
}

type XORResponseBody struct {
	Message string `json:"message"`
}

func CipherXOR(w http.ResponseWriter, r *http.Request) {
	var requestBody XORRequestBody
	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var cipherText string
	var err error
	switch strings.ToLower(requestBody.Type) {
	case "encrypt":
		cipherText = ciphers.EncryptXOR(requestBody.Message, requestBody.Key)
	case "decrypt":
		cipherText, err = ciphers.DecryptXOR(requestBody.Message, requestBody.Key)
	default:
		writeError(w, http.StatusBadRequest, "Invalid type: must be 'encrypt' or 'decrypt'")
		return
	}

	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseBody := XORResponseBody{
		Message: cipherText,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}