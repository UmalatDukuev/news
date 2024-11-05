package handler

import (
	"net/http"
	"news/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() {
	http.HandleFunc("/", h.homePage)
}

func (h *Handler) homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running..."))
}
