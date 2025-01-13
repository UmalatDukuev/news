package handler

import (
	"github.com/UmalatDukuev/news/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

	}

	posts := router.Group("/posts")
	{
		posts.POST("/", h.createPost)
		posts.GET("/", h.getAllPosts)
		posts.GET("/:id", h.getPostById)
		likes := posts.Group("/:id/likes", h.userIdentity)
		{
			likes.POST("/", h.createLike)
		}

		comments := posts.Group("/:id/comments", h.userIdentity)
		{
			comments.POST("/", h.createComment)
		}

		tags := posts.Group("/:id/tags", h.userIdentity)
		{
			tags.POST("/", h.createTag)
		}

	}

	return router
}
