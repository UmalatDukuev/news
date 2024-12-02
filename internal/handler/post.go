package handler

import (
	"net/http"
	"news/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) homePage(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Server is running..."})
}

func (h *Handler) createPost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := h.service.PostService.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}
