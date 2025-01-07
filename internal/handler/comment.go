package handler

import (
	"fmt"

	"github.com/UmalatDukuev/news/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createComment(c *gin.Context) {
	postID := c.Param("id")
	var comment models.Comment
	fmt.Print("ID: ")
	fmt.Println(postID)
	_, _ = h.services.Comment.Create(comment)

}
