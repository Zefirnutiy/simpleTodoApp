package repository

import (
	"Zefirnutiy/simpleTodoApp/structs"

	"github.com/jmoiron/sqlx"
)

type Authorization interface{
	CreateUser(user structs.User)(int, error)
	GetUser(userName, password string)(structs.User, error)
}

type Todo interface{

}

type TodoList interface{

}

type Repository struct{
	Authorization
	Todo
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}