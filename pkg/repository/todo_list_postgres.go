package repository

import (
	"Zefirnutiy/simpleTodoApp/structs"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres{
	return &TodoListPostgres{db:db}
}

func (r *TodoListPostgres)Create(userId int, list structs.TodoList)(int, error){
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	var listId int
	
	createListQuery := fmt.Sprintf("INSERT INTO %s ( title, description ) VALUES( $1, $2 ) RETURNING id", todoListTable)
	row := r.db.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES( $1, $2 )", usersListsTable)
	_, err = r.db.Exec(createUsersListQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return listId, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int)([]structs.TodoList, error){
	var todoList []structs.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
	todoListTable, usersListsTable)

	err := r.db.Select(&todoList, query, userId)

	return todoList, err
}

func (r *TodoListPostgres) GetListById(listId, userId int)(structs.TodoList, error){
	var todoList structs.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
						INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id=$2`,
		todoListTable, usersListsTable)
	err := r.db.Get(&todoList, query, userId, listId)

	return todoList, err
}

func (r *TodoListPostgres) DeleteList(userId, listId int)error{
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id=ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", 
	todoListTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *TodoListPostgres) Update(userId, listId int, input structs.UpdateTodoList) error{
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValue = append(setValue, fmt.Sprintf("title=$%d", argId))
		args = append(args, &input.Title)
		argId++
	}

	if input.Description != nil {
		setValue = append(setValue, fmt.Sprintf("description=$%d", argId))
		args = append(args, &input.Description)
		argId++
	}

	setQuery := strings.Join(setValue, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debug("updateQuery: %s", query)
	logrus.Debug("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}