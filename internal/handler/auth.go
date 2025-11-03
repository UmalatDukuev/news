package handler

import (
	"net/http"

	"github.com/UmalatDukuev/news/internal/errs"
	"github.com/UmalatDukuev/news/internal/utils"
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
		er := errs.ErrInvalidJSON
		newErrorResponse(c, errs.ErrorCodeToHTTPStatus[er], er.Error())
		return
	}

	if len(input.Username) < 3 || len(input.Password) < 3 {
		er := errs.ErrInvalidCredentials
		newErrorResponse(c, errs.ErrorCodeToHTTPStatus[er], er.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(models.User{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		c.JSON(errs.ErrorCodeToHTTPStatus[err], gin.H{
			"error": err.Error(),
		})
		return
	}
	tokenPair, err := utils.GenerateToken(id)
	utils.SetAuthCookies(c.Writer, tokenPair)
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.services.Authorization.SignIn(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, errs.ErrorCodeToHTTPStatus[err], err.Error())
		return
	}

	tokenPair, err := utils.GenerateToken(user.ID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to generate tokens")
		return
	}

	utils.SetAuthCookies(c.Writer, tokenPair)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
