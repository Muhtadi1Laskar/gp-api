package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
	"strings"
)

type CeaserRequestBody struct {
	Message string `json:"message" validate:"required"`
	Key     int    `json:"key" validate:"required"`
	Type    string `json:"type" validate:"required"`
}

type CeaserResponseBody struct {
	Message string `json:"message"`
}

func CeaserCipher(w http.ResponseWriter, r *http.Request) {
	var requestBody CeaserRequestBody
	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	var cipher string
	switch strings.ToLower(requestBody.Type) {
	case "encrypt":
		cipher = ciphers.EncryptCaeser(requestBody.Message, requestBody.Key)
	case "decrypt":
		cipher = ciphers.EncryptCaeser(requestBody.Message, -requestBody.Key)
	default:
		writeError(w, http.StatusBadRequest, "Invalid type: must be 'encrypt' or 'decrypt'")
		return
	}

	responseBody := CeaserResponseBody{
		Message: cipher,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}