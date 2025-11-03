package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/UmalatDukuev/news/internal/errs"
	"github.com/UmalatDukuev/news/internal/utils"
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

func (r *AuthPostgres) GetUser(username string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id, username, password_hash FROM %s WHERE username=$1", usersTable)
	err := r.db.Get(&user, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errs.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}

func (r *AuthPostgres) CheckPassword(username, password string) (bool, error) {
	user, err := r.GetUser(username)
	if err != nil {
		return false, err
	}
	isValid := utils.CheckPasswordHash(password, user.Password)
	if !isValid {
		return false, errs.ErrInvalidCredentials
	}
	return true, nil
}
