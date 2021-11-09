package controllers

import (
	"go-todo-app/config"
	"net/http"
)

func StartMainServer() error {
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/authenticate", authenticate)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
