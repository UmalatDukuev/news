package app

import (
	"log"
	"net/http"
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
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	db, err := utils.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	router := router()
	handler.InitRoutes(router)

	router.Run(":8000")

	log.Println("Server is running on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}
}

func router() *gin.Engine {
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
