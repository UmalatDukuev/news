package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/UmalatDukuev/news/internal/errs"
	"github.com/UmalatDukuev/news/models"
	"github.com/jmoiron/sqlx"

	"github.com/lib/pq"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (username, password_hash) values ($1, $2) RETURNING id",
		usersTable,
	)

	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return 0, errs.ErrUsernameTaken

			}
		}
		return 0, err
	}

	return id, nil
}

// func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
// 	var user models.User
// 	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
// 	err := r.db.Get(&user, query, username, password)

// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return user, errs.ErrUserNotFound
// 		}
// 		return user, err
// 	}

// 	return user, nil
// }

func (r *AuthPostgres) GetUser(username string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1", usersTable)
	err := r.db.Get(&user, query, username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errs.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}

func (r *AuthPostgres) СheckPassword(username, passwordHash string) (bool, error) {
	query := fmt.Sprintf("select password_hash from %s where username = $1", usersTable)
	passwordH := ""
	err := r.db.Get(&passwordH, query, username)
	if errors.Is(err, sql.ErrNoRows) {
		return false, errs.ErrUserNotFound
	}

	if passwordH == passwordHash {
		return true, nil
	}
	return false, nil
}
