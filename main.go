package main

import (
	"fmt"
	"go-todo-app/app/models"
)

func main() {
	//fmt.Println(config.Config.Port)
	//fmt.Println(config.Config.SQLDriver)
	//fmt.Println(config.Config.DbName)
	//fmt.Println(config.Config.LogFile)
	//
	//log.Println("test")
	//fmt.Println(models.Db)
	//
	//u := &models.User{}
	//u.Name = "Hiroshi"
	//u.Email = "hiroshi@example.com"
	//u.PassWord = "testtest"
	//fmt.Println(u)

	//u.CreateUser()
	//
	//user, _ := models.GetUser(2)
	//user.CreateTodo("CD返却")

	//fmt.Printf("%+v\n", u)
	//
	//u.Name = "Takeshi Emoto"
	//u.Email = "test3@example.com"
	//u.UpdateUser()
	//u, _ = models.GetUser(2)
	//fmt.Printf("%+v\n", u)
	//u.DeleteUser()

	//t, _ := models.GetTodo(1)
	//fmt.Println(t)

	//todos, _ := models.GetTodos()
	//fmt.Println(todos)

	//u2, _ := models.GetUser(2)
	//todos, _ := u2.GetTodosByUser()
	//for _, v := range todos {
	//	fmt.Println(v)
	//}

	t, _ := models.GetTodo(1)
	t.Content = "Hello world"
	err := t.UpdateTodo()
	if err != nil {
		fmt.Println(err)
	}

	t.DeleteTodo()

}
