package handler

import (
	"github.com/UmalatDukuev/news/internal/errs"
	"github.com/UmalatDukuev/news/internal/utils"
	"github.com/gin-gonic/gin"
)

const userCtx = "userId"

func (h *Handler) userIdentity(c *gin.Context) {
	accessCookie, err := c.Request.Cookie("access_token")

	if err != nil {
		//...........
	}

	userID, err := utils.ValidateAccessToken(accessCookie.Value)
	if err != nil {
		handleRefreshFlow(c)
		return
	}

	c.Set(userCtx, userID)
	c.Next()
}

func handleRefreshFlow(c *gin.Context) {
	refreshCookie, err := c.Request.Cookie("refresh_token")
	if err != nil {
		newErrorResponse(c, errs.ErrorCodeToHTTPStatus[errs.ErrMissingTokens], errs.ErrMissingTokens.Error())
		return
	}

	userID, err := utils.ValidateRefreshToken(refreshCookie.Value)
	if err != nil {
		newErrorResponse(c, errs.ErrorCodeToHTTPStatus[errs.ErrInvalidToken], errs.ErrInvalidToken.Error())
		return
	}

	tokens, err := utils.GenerateToken(userID)
	if err != nil {
		newErrorResponse(c, errs.ErrorCodeToHTTPStatus[errs.ErrTokenGeneration], errs.ErrTokenGeneration.Error())
		return
	}

	utils.SetAuthCookies(c.Writer, tokens)
	c.Set(userCtx, userID)
	c.Next()
}
