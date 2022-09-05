package repository

import (
	"Zefirnutiy/simpleTodoApp/structs"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres{
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres ) CreateUser(user structs.User)(int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash ) values($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.UserName, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(userName, password string)(structs.User, error){
	var user structs.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, userName, password)

	return user, err
}