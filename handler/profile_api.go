package handler

import (
	"google-oauth/model"
	"google-oauth/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfileApi(c *gin.Context) {

	user, exists := c.Get("user")
	if !exists {
		c.JSON(404, gin.H{"error": "value not found"})
		return
	}
	authUser := user.(model.User)

	userResponse := web.UserResponse{
		Email:   authUser.Email,
		Name:    authUser.Name,
		Picture: authUser.Picture,
	}

	c.JSON(http.StatusOK, userResponse)
}
