package handler

import (
	"fmt"

	"github.com/UmalatDukuev/news/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createTag(c *gin.Context) {
	postID := c.Param("id")
	var tag models.Tag
	fmt.Print("ID: ")
	fmt.Println(postID)
	_, _ = h.services.Tag.Create(tag)

}
