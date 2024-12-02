package repository

import (
	"news/internal/handler/request"

	"github.com/gin-gonic/gin"
)

func (r *PostRepository) Auth(c *gin.Context, req request.Login) error {
	return nil
}
