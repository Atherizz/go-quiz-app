package handler

import (
	"google-oauth/helper"
	"google-oauth/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {

	user, exists := c.Get("user")
	if !exists {
		c.JSON(404, gin.H{"error": "value not found"})
		return
	}
	authUser := user.(helper.UserSession)

	userResponse := web.UserResponse{
		Email:   authUser.Email,
		Name:    authUser.Name,
		Picture: authUser.Picture,
	}
	

	c.JSON(http.StatusOK, userResponse)
}
