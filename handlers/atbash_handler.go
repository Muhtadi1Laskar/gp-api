package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
)


func AtbashHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody KeyLessCipherRequest

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return 
	}

	cipherText := ciphers.AtbashCipher(requestBody.Message)

	responseBody := KeyLessCipherResponse{
		Message: cipherText,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}