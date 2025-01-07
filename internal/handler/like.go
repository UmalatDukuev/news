package handler

import (
	"fmt"

	"github.com/UmalatDukuev/news/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createLike(c *gin.Context) {
	postID := c.Param("id")
	var like models.Like
	fmt.Print("ID: ")
	fmt.Println(postID)
	_, _ = h.services.Like.Create(like)

}
