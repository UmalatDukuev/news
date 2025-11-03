package utils

import (
	"errors"

	errs "github.com/UmalatDukuev/news/internal/errs"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		return "", errs.ErrHashingPass
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}
