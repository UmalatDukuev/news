package handler

import (
	"fmt"
	"news/internal/handler/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) login(c *gin.Context) {
	var req request.Login
	_ = c.ShouldBindJSON(&req)
	fmt.Println(req)
}
