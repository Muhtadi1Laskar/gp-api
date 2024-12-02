package handlers

import (
	"fmt"
	hashes "go-api/Hashes"
	"io"
	"net/http"
)

func HandleFileHash(w http.ResponseWriter, r *http.Request) {
	fileText, err := UploadFile(r)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	hash, err := hashes.HashData(fileText["message"], fileText["hash"])
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseBody := KeyLessCipherResponse{
		Message: hash,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
} 



func UploadFile(r *http.Request) (map[string]string, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, fmt.Errorf("unable to read request body: %v", err)
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		return nil, fmt.Errorf("unable to read request body: %v", err)
	}
	defer file.Close()


	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read file bytes %v", err)
	}

	temp := map[string]string{
		"message": string(fileBytes),
	}

	for key, data := range r.MultipartForm.Value {
		temp[key] = data[0]
	}

	return temp, nil
}
