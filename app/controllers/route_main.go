package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string
}

func top(w http.ResponseWriter, r *http.Request) {
	msg := Message{Text: "Hello world"}
	res, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}
