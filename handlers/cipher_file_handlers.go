package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
	"strings"
)

func HandleCipherFile(w http.ResponseWriter, r *http.Request) {
	requestBody, err := UploadFile(r)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var cipherText string

	switch strings.ToLower(requestBody["type"]) {
	case "encrypt":
		cipherText, err = ciphers.AESEncrypt(requestBody["message"], requestBody["key"])
	case "decrypt":
		cipherText, err = ciphers.AESDecrypt(requestBody["message"], requestBody["key"])
	default:
		writeError(w, http.StatusBadRequest, "Invalid type: must be 'encrypt' or 'decrypt'")
		return
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