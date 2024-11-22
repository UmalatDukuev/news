package app

import (
	"log"
	"news/internal/handler"
	"news/internal/repository"
	"news/internal/service"
	"news/internal/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	cfg, err := repository.LoadDBConfig()
	if err != nil {
		log.Fatalf("Failed to load DB config: %s", err.Error())
	}

	db, err := utils.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	log.Println("Database connection established.")

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	router := initRouter()
	handler.InitRoutes(router)

	log.Println("Server is running on port 8000")
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}

func initRouter() *gin.Engine {
	var r *gin.Engine

	if env := os.Getenv("APP_ENV"); env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery())
	} else {
		r = gin.Default()
	}

	return r
}
