package main

import (
	"log"

	"github.com/Zhandospa/todo-app"
	"github.com/Zhandospa/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occoured wjile running http server: %s", err.Error())
	}
}
