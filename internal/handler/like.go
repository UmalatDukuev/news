package handler

import (
	"strconv"

	"github.com/UmalatDukuev/news/internal/utils"
	"github.com/UmalatDukuev/news/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createLike(c *gin.Context) {
	postID := c.Param("id")
	userID, _ := utils.GetUserId(c)
	var like models.Like
	like.PostID, _ = strconv.Atoi(postID)
	like.UserID = userID
	_, _ = h.services.Like.Create(like)

}

func (h *Handler) getAllPostLikes(c *gin.Context) {

	_, err := h.services.GetAllPostLikes(1)
	if err != nil {

	}

}
