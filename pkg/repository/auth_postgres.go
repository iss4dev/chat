package repository

import (
	"fmt"
	todo "to-do-list"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash, role) VALUES ($1, $2, $3, $4) RETURNING id",
		usersTable,
	)

	err = tx.QueryRow(query, user.Name, user.Username, user.Password, user.Role).Scan(&id)
	if err != nil {
		return 0, err
	}

	profileQuery := fmt.Sprintf("INSERT INTO %s (user_id, avatar_url, bio, updated_at) VALUES ($1, '', '', CURRENT_TIMESTAMP)", userProfilesTable)

	_, err = tx.Exec(profileQuery, id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id, name, username, password_hash, role FROM %s WHERE username=$1", usersTable)

	err := r.db.Get(&user, query, username)
	if err != nil {
		return todo.User{}, err
	}

	return user, nil
}
