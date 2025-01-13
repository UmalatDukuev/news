package handler

import (
	"net/http"
	"strconv"

	"github.com/UmalatDukuev/news/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createComment(c *gin.Context) {
	postID := c.Param("id")
	userID, _ := getUserId(c)
	var comment models.Comment
	comment.PostID, _ = strconv.Atoi(postID)
	comment.UserID = userID
	if err := c.BindJSON(&comment); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	_, _ = h.services.Comment.Create(comment)

}
