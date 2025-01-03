package handler

import (
	"net/http"

	"github.com/UmalatDukuev/news"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	var input news.Post

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	id, err := h.services.Post.CreatePost(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
