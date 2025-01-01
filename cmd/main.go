package main

import (
	"log"

	"github.com/UmalatDukuev/news"
	"github.com/UmalatDukuev/news/internal/handler"
	"github.com/UmalatDukuev/news/internal/repository"
	"github.com/UmalatDukuev/news/internal/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(news.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Printf("%s", &err)
	}
}
