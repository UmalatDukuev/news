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

func (h *Handler) InitRoutes(router *gin.Engine) {
	// http.HandleFunc("/", h.homePage)
	// v1 := router.Group("/auth")
	// {
	// 	v1.GET("/", h.homePage)

	// 	posts := v1.Group("/posts")
	// 	{
	// 		posts.POST("/", h.CreatePost)
	// 		// 	// posts.GET("/", h.GetPosts)
	// 		// 	// posts.GET("/:id", h.GetPost)
	// 		// 	// posts.PUT("/:id", h.UpdatePost)
	// 		// 	// posts.DELETE("/:id", h.DeletePost)
	// 	}
	// }
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
		// 	// posts.GET("/", h.GetPosts)
		// 	// posts.GET("/:id", h.GetPost)
		// 	// posts.PUT("/:id", h.UpdatePost)
		// 	// posts.DELETE("/:id", h.DeletePost)
	}

	v1 := router.Group("/v1")
	{
		v1.GET("/", h.homePage)
		posts := v1.Group("/posts")
		{
			posts.POST("/", h.createPost)
			// 	// posts.GET("/", h.GetPosts)
			// 	// posts.GET("/:id", h.GetPost)
			// 	// posts.PUT("/:id", h.UpdatePost)
			// 	// posts.DELETE("/:id", h.DeletePost)
		}
	}
}
