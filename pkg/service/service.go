package service

import (
	"Zefirnutiy/simpleTodoApp/pkg/repository"
	"Zefirnutiy/simpleTodoApp/structs"
)

type Authorization interface{
	CreateUser(user structs.User)(int, error)
	GenerateToken(userName, password string)(string, error)
	ParseToken(token string)(int, error)
}


type TodoList interface{
	Create(userId int, list structs.TodoList)(int, error)
	GetAll(userId int)([]structs.TodoList, error)
	GetListById(listId, userId int)(structs.TodoList, error)
	DeleteList(userId, listId int)error
	Update(userId, listId int, input structs.UpdateTodoList) error
}
type Todo interface{
	Create(userId, listId int, todo structs.Todo)(int, error)
	GetAll(userId, listId int)([]structs.Todo, error)
	GetTodoById(todoId, userId int)(structs.Todo, error)
	Delete(userId, todoId  int)(error)
	Update(userId, todoId int, input structs.UpdateTodo) error
}

type Service struct{
	Authorization
	TodoList
	Todo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		TodoList: NewTodoListService(repos.TodoList),
		Todo: NewTodoService(repos.Todo, repos.TodoList),
	}
}