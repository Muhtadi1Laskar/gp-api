package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
	"strings"
)

type VigenereRequestBody struct {
	Message string `json:"message" validate:"required"`
	Key     string    `json:"key" validate:"required"`
	Type    string `json:"type" validate:"required"`
}

type VingereResponseBody struct {
	Message string `json:"message"`
}

func VigenereCipherHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody VigenereRequestBody
	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	var cipher string
	switch strings.ToLower(requestBody.Type) {
	case "encrypt":
		cipher = ciphers.VigenereCipher(requestBody.Message, requestBody.Key, true)
	case "decrypt":
		cipher = ciphers.VigenereCipher(requestBody.Message, requestBody.Key, false)
	default:
		writeError(w, http.StatusBadRequest, "Invalid type: must be 'encrypt' or 'decrypt'")
		return
	}

	responseBody := CeaserResponseBody{
		Message: cipher,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}