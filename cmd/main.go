package main

import (
	"log"
	"net/http"
	"news/internal/handler"
	"news/internal/repository"
	"news/internal/service"
)

func main() {
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	handler.InitRoutes()

	// Запускаем сервер
	log.Println("Server is running on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}
}
