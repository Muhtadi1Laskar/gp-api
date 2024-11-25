package handler

import (
	"encoding/json"
	"fmt"
	hashes "go-api/Hashes"
	"io"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
	Hash string `json:"hash"`
}

type ResponseBody struct {
	Hash string `json:"hash"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedData, err := hashes.HashData(requestBody.Message, requestBody.Hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ResponseBody{
		Hash: hashedData,
	}

	writeJSONResponse(w, http.StatusOK, response)
}

func readRequestBody(r *http.Request, target interface{}) error {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("unable to read request body: %v", err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(reqBody, target); err != nil {
		return fmt.Errorf("unable to read request body: %v", err)
	}
	return nil
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func Main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
