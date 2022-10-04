package repository

import (
	"Zefirnutiy/simpleTodoApp/structs"

	"github.com/jmoiron/sqlx"
)

type Authorization interface{
	CreateUser(user structs.User)(int, error)
	GetUser(userName, password string)(structs.User, error)
}


type TodoList interface{
	Create(userId int, list structs.TodoList)(int, error)
	GetAll(userId int)([]structs.TodoList, error)
	GetListById(listId, userId int)(structs.TodoList, error)
	DeleteList(userId, listId int) error
	Update(userId, listId int, input structs.UpdateTodoList) error
}

type Todo interface{
	Create(listId int, todo structs.Todo) (int, error)
	GetAll(userId, listId int)([]structs.Todo, error)
	GetTodoById(todoId, userId int)(structs.Todo, error)
	Delete(todoId, userId int) error
	Update(userId, todoId int, input structs.UpdateTodo) error
}
type Repository struct{
	Authorization
	TodoList
	Todo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
		Todo: NewTodoPostgres(db),
	}
}