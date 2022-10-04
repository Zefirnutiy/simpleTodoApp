package repository

import (
	"Zefirnutiy/simpleTodoApp/structs"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type TodoPostgres struct {
	db *sqlx.DB
}

func NewTodoPostgres(db *sqlx.DB) *TodoPostgres{
	return &TodoPostgres{db: db}
}

func (r *TodoPostgres) Create(listId int, todo structs.Todo) (int, error){
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createTodoQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES($1, $2) RETURNING id", todoTable)
	row := tx.QueryRow(createTodoQuery, todo.Title, todo.Description)
	
	var todoId int

	if err := row.Scan(&todoId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListTodoQuery := fmt.Sprintf("INSERT INTO %s (list_id, todo_id ) VALUES($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListTodoQuery, listId, todoId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return todoId, nil
}

func (r *TodoPostgres) GetAll(userId, listId int)([]structs.Todo, error){
	var todos []structs.Todo

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description FROM %s ti INNER JOIN %s li on li.todo_id = ti.id
						INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
						todoTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&todos, query, listId, userId); err != nil {
		return nil, err
	}
	
	return todos, nil
}

func (r *TodoPostgres) GetTodoById(todoId, userId int)(structs.Todo, error){
	var todo structs.Todo

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description FROM %s ti INNER JOIN %s li on li.todo_id = ti.id
						INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
						todoTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&todo, query, todoId, userId); err != nil {
		return todo, err
	}
	
	return todo, nil
}

func (r *TodoPostgres) Delete(todoId, userId int)(error){

	query := fmt.Sprintf(`DELETE FROM %s tt USING %s tl, %s ul 
						WHERE tt.id = tl.todo_id AND tl.list_id = ul.list_id AND ul.user_id = $1 AND tt.id = $2`,
					todoTable, listsItemsTable, usersListsTable)

	_, err := r.db.Exec(query, userId, todoId)

	return err
}

func (r *TodoPostgres) Update(userId, todoId int, input structs.UpdateTodo)(error){
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValue = append(setValue, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValue = append(setValue, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValue = append(setValue, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValue, ", ")

	fmt.Println(setQuery)
	query := fmt.Sprintf(`UPDATE %s td SET %s FROM %s lt ,%s ul 
				WHERE td.id = lt.todo_id AND lt.list_id = ul.list_id AND ul.user_id=$%d AND td.id=$%d`,
		todoTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, todoId)

	_, err := r.db.Exec(query, args...)
	return err
}