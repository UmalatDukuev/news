package handler

import (
	"news/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes(router *gin.Engine, service *service.Service) {
	// http.HandleFunc("/", h.homePage)
	// http.HandleFunc("/", h.homePage)
	v1 := router.Group("/v1")
	{
		// Эндпоинт для главной страницы
		v1.GET("/", h.homePage)

		posts := v1.Group("/posts")
		{
			posts.POST("/", h.CreatePost) // Создать пост
			// 	// posts.GET("/", h.GetPosts)         // Получить список постов
			// 	// posts.GET("/:id", h.GetPost)       // Получить пост по ID
			// 	// posts.PUT("/:id", h.UpdatePost)    // Обновить пост
			// 	// posts.DELETE("/:id", h.DeletePost) // Удалить пост
		}
	}
}
