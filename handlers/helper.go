package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
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

func writeError(w http.ResponseWriter, statusCode int, err string) {
	writeJSONResponse(w, statusCode, ErrorResponse{ Error: err })
}
