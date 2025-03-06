package handler

import (
	"net/http"
	"strconv"

	"github.com/UmalatDukuev/news/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	var input models.Post
	userID, _ := getUserId(c)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	input.AuthorID = userID
	id, err := h.services.Post.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updatePost(c *gin.Context) {
	var input models.Post
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	userID, _ := getUserId(c)
	input.AuthorID = userID

	input.ID = postID
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if input.Description == "" && input.Title == "" {
		newErrorResponse(c, http.StatusBadRequest, "title and description can't be both empty")
		return
	}
	err = h.services.Post.Update(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": postID,
	})
}

func (h *Handler) getAllPosts(c *gin.Context) {
	posts, err := h.services.Post.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *Handler) getPostById(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	post, err := h.services.Post.GetById(postID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, post)
}
