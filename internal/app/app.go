package app

import (
	"log"
	"net/http"
	"news/internal/handler"
	"news/internal/repository"
	"news/internal/service"
	"news/internal/utils"
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

	handler.InitRoutes()

	log.Println("Server is running on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}
}
