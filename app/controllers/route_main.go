package controllers

import (
	"encoding/json"
	"fmt"
	"go-todo-app/app/dto"
	"go-todo-app/app/models"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	indexDto := dto.IndexDto{Message: "Logged In"}

	j, err := json.Marshal(indexDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(j))

	return
}

func todos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	s, err := session(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	u, err := s.GetUserBySession()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	var todos = []models.Todo{}

	t, err := u.GetTodosByUser()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	if t != nil {
		todos = t
	}

	todosResponseDto := dto.TodosResponseDto{
		Data: todos,
	}

	j, err := json.Marshal(todosResponseDto)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(j))

	return
}
