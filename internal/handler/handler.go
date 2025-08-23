package handler

import (
	"time"

	"github.com/UmalatDukuev/news/internal/service"
	"github.com/gin-contrib/cors"
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

	cfg := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(cfg))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

	}

	posts := router.Group("/posts", h.userIdentity)
	{
		posts.POST("/", h.createPost)
		posts.GET("/", h.getAllPosts)
		posts.GET("/:id", h.getPostById)
		posts.PUT("/:id", h.updatePost)

		likes := posts.Group("/:id/likes")
		{
			likes.GET("/", h.getAllPostLikes)
			likes.POST("/", h.createLike)
		}

		comments := posts.Group("/:id/comments")
		{
			comments.POST("/", h.createComment)
		}

		tags := posts.Group("/:id/tags")
		{
			tags.POST("/", h.createTag)
		}

	}

	return router
}
