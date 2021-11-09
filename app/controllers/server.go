package controllers

import (
	"fmt"
	"go-todo-app/app/models"
	"go-todo-app/config"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (session models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		session = models.Session{UUID: cookie.Value}
		if ok, _ := session.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}

	return session, err
}

func StartMainServer() error {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/signout", signout)

	http.HandleFunc("/todos", todos)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
