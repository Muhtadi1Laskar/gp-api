package handlers

import (
	"go-api/Encoding"
	"net/http"
	"strings"
)

type BaseEncoderRequest struct {
	Message     string `json:"message" validate:"required"`
	EncoderType string `json:"encoderType" validate:"required"`
	Type        string `json:"type" validate:"required"`
}


func BaseEncoderHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody BaseEncoderRequest

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var encodeType string = strings.ToLower(requestBody.EncoderType)
	var encodedString string
	var err error
	switch strings.ToLower(requestBody.Type) {
	case "encode":
		encodedString, err = Encoding.EncodeToBase(requestBody.Message, encodeType)
	case "decode":
		encodedString, err = Encoding.DecodeFromBase(requestBody.Message, encodeType)
	default:
		writeError(w, http.StatusBadRequest, "Invalid type: must be 'encode' or 'decode'")
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