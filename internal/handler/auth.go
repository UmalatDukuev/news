package handler

import (
	"fmt"
	"net/http"
	"news/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) login(c *gin.Context) {

}

func (h *Handler) signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if 
	fmt.Println(user.Password)
}
