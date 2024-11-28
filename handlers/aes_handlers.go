package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
	"strings"
)

type AESRequestBody struct {
	Message string `json:"message" validate:"required"`
	Key string `json:"key" validate:"required"`
	Type string `json:"type" validate: "required"`
}

type AESResponseBody struct {
	Message string `json:"message"`
}

func CipherAES(w http.ResponseWriter, r *http.Request) {
	var requestBody AESRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	var cipherText string
	var err error
	switch strings.ToLower(requestBody.Type) {
	case "encrypt":
		cipherText, err = ciphers.AESEncrypt(requestBody.Message, requestBody.Key)
	case "decrypt":
		cipherText, err = ciphers.AESDecrypt(requestBody.Message, requestBody.Key)
	default:
		writeError(w, http.StatusBadRequest, "Invalid type: must be 'encrypt' or 'decrypt'")
	}

	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseBody := AESResponseBody{
		Message: cipherText,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}