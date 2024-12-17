package repository

import (
	"fmt"
	"news/internal/handler/request"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *PostRepository {
	if db == nil {
		fmt.Println("!")
	}
	fmt.Println(".")
	return &PostRepository{db: db}
}

func (r *PostRepository) Auth(c *gin.Context, req request.Login) error {
	return nil
}
