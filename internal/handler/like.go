package handler

import (
	"fmt"
	"strconv"

	"github.com/UmalatDukuev/news/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createLike(c *gin.Context) {
	postID := c.Param("id")
	userID, _ := getUserId(c)
	var like models.Like
	like.PostID, _ = strconv.Atoi(postID)
	like.UserID = userID
	_, _ = h.services.Like.Create(like)

}

func (h *Handler) getAllPostLikes(c *gin.Context) {
	// postID := c.Param("id")
	// var likes []models.LikeOnPost
	// fmt.Println(5555555)
	// fmt.Println(postID, likes)

	// like.PostID, _ = strconv.Atoi(postID)
	_, err := h.services.GetAllPostLikes(1)
	if err != nil {
		fmt.Println(123)
	}
	// _, _ = h.services.Like.Create(like)

}
