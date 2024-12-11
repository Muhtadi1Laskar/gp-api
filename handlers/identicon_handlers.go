package handlers

import (
	"go-api/Encoding"
	"net/http"
)

type IdenticonResponse struct {
	Identicon string
}

func IdenticonHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody HashRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	identiconImageString, err := Encoding.IdenticonMain(requestBody.Message, requestBody.Hash)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}


	responseBody := IdenticonResponse{
		Identicon: identiconImageString,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}