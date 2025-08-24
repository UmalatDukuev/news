package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/UmalatDukuev/news/internal/errs"

	"github.com/UmalatDukuev/news/models"
	"github.com/gin-gonic/gin"
)

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input signInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid JSON body")
		return
	}

	input.Username = strings.TrimSpace(input.Username)
	if input.Username == "" || len(input.Password) < 3 {
		newErrorResponse(c, http.StatusBadRequest, "username or password is invalid")
		return
	}

	id, err := h.services.Authorization.CreateUser(models.User{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		if status, ok := errs.ErrorCodeToHTTPStatus[err]; ok {
			newErrorResponse(c, status, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
		// "token":
	})
}

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrUserNotFound):
			newErrorResponse(c, http.StatusBadRequest, "invalid username or password")
		case errors.Is(err, errs.ErrInvalidCredentials):
			newErrorResponse(c, http.StatusBadRequest, "invalid username or password")

		default:
			newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
