package main

import (
	"Zefirnutiy/simpleTodoApp"
	"Zefirnutiy/simpleTodoApp/pkg/handler"
	"log"
)

func main(){
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil{
		log.Fatalf(`Ошибка запуска сервера: %s`, err.Error())
	}
}