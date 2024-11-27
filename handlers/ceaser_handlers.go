package handlers

import (
	ciphers "go-api/Ciphers"
	"net/http"
)

type CeaserRequestBody struct {
	Message string `json:"message" validate:"required"` 
	Key int `json:"int" validate:"required"`
}

type CeaserResponseBody struct {
	Message string `json:"message"`
}

func CeaserCipher(w http.ResponseWriter, r *http.Request) {
	var requestBody CeaserRequestBody
	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	cipher := ciphers.CeaserCipher(requestBody.Message, requestBody.Key)

	responseBody := CeaserResponseBody{
		Message: cipher,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}