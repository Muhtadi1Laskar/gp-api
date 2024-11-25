package handler

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	ID int `json:"id"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	person := Person{ Name: "Gino", ID: 343242543242 }

	jsonStr, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonStr)
}
