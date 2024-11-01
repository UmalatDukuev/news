package main

import (
	"fmt"
	"net/http"
	"news/internal/handler"
	"news/internal/repository"
	"news/internal/service"
)

func main() {
	repo := repository.NewRepository()
	service := service.NewService()
	handler := handler.NewHandler()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is running...")
	})
	http.ListenAndServe(":8000", nil)
	return
}
