package controllers

import (
	"encoding/json"
	"fmt"
	"go-todo-app/app/dto"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	indexDto := dto.IndexDto{Message: "Logged In"}

	res, err := json.Marshal(indexDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))

	return
}
