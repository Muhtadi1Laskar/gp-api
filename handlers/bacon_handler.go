package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
	"strings"
)

func BaconCipherHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody KeyLessCipherRequest

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var cipherText string
	var err error

	switch strings.ToLower(requestBody.Type) {
	case "encrypt":
		cipherText = ciphers.BaconCipherEncrypt(requestBody.Message)
	case "decrypt":
		cipherText, err = ciphers.BaconCipherDecrypt(requestBody.Message)
	default:
		writeError(w, http.StatusBadRequest, "Invalid type: must be 'encrypt' or 'decrypt'")
		return
	}

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseBody := KeyLessCipherResponse{
		Message: cipherText,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}