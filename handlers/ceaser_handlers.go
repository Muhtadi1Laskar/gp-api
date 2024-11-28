package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
	"strconv"
	"strings"
)

type CeaserRequestBody struct {
	Message string `json:"message" validate:"required"`
	Key     string    `json:"key" validate:"required"`
	Type    string `json:"type" validate:"required"`
}

type CeaserResponseBody struct {
	Message string `json:"message"`
}

func CeaserCipherHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody CeaserRequestBody
	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	intKey, _ := strconv.Atoi(requestBody.Key)

	var cipher string
	switch strings.ToLower(requestBody.Type) {
	case "encrypt":
		cipher = ciphers.EncryptCaeser(requestBody.Message, intKey)
	case "decrypt":
		cipher = ciphers.EncryptCaeser(requestBody.Message, -intKey)
	default:
		writeError(w, http.StatusBadRequest, "Invalid type: must be 'encrypt' or 'decrypt'")
		return
	}

	responseBody := CeaserResponseBody{
		Message: cipher,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}