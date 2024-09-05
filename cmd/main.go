package main

import (
	"log"

	"github.com/Zhandospa/todo-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatal("error occoured wjile running http server: %s", err.Error())
	}
}
