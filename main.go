package main

import (
	"fmt"
	"go-todo-app/app/controllers"
	"go-todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
