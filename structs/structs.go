package structs

import "errors"

type User struct {
	Id 		 int 	`json:"-" db:"id"`
	Name 	 string `json:"name" binding:"required"`
	UserName string `json:"userName" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

type Todo struct {
	Id 		 	int 	`json:"id" db:"id"`
	Title 	 	string  `json:"title" db:"title" binding:"required"`
	Description string  `json:"description" db:"description"`
	Done 		bool	`json:"done" db:"done"`
}

type TodoList struct {
	Id 		 	int 	`json:"id" db:"id"`
	Title 	 	string  `json:"title" binding:"required" db:"title"`
	Description string  `json:"description" db:"description"`
}

type UsersLists struct {
	Id 		int	`json:"id"`
	UserID	int	`json:"userId"`
	ListID	int	`json:"listId"`
}

type ListsTodos struct {
	Id 		int	`json:"id"`
	ListID	int	`json:"listId"`
	TodoID	int	`json:"todoId"`
}

type UpdateTodoList struct {
	Title *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateTodoList) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("updates structure has no values")
	}

	return nil
}

type UpdateTodo struct {
	Title 		*string `json:"title"`
	Description *string `json:"description"`
	Done 		*bool	`json:"done"`
}

func (i UpdateTodo) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil{
		return errors.New("updates structure has no values")
	}

	return nil
}