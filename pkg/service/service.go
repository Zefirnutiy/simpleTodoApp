package service

import (
	"Zefirnutiy/simpleTodoApp/pkg/repository"
	"Zefirnutiy/simpleTodoApp/structs"
)

type Authorization interface{
	CreateUser(user structs.User)(int, error)

}

type Todo interface{

}

type TodoList interface{

}

type Service struct{
	Authorization
	Todo
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
	}
}