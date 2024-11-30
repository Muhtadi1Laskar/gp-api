package handlers

import (
	"go-api/Encoding"
	"net/http"
	"strings"
)

func BinaryEncoderHander(w http.ResponseWriter, r *http.Request) {
	var requestBody KeyLessCipherRequest

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var encodedString string
	var err error
	switch strings.ToLower(requestBody.Type) {
	case "encrypt": 
		encodedString = Encoding.TextToBinary(requestBody.Message)
	case "decrypt":
		encodedString, err = Encoding.BinaryToText(requestBody.Type)
	default:
		writeError(w, http.StatusBadRequest, "Invalid type: must be 'encrypt' or 'decrypt'")
		return 
	}

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return 
	}

	responseBody := KeyLessCipherResponse{
		Message: encodedString,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}