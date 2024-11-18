package handler

import (
	"net/http"
	"news/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) homePage(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Server is running..."})
}

func (h *Handler) CreatePost(c *gin.Context) {
	var post models.Post

	// Пробуем привязать данные из тела запроса к модели Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Вызов сервиса для создания поста
	if err := h.service.PostService.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем успешный ответ с созданным постом
	c.JSON(http.StatusOK, post)
}
