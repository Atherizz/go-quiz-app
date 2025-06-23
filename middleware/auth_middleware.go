package middleware

import (
	"google-oauth/helper"
	"google-oauth/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := helper.Store.Get(c.Request, "user_info")

		user, ok := session.Values["user"].(model.User)
		if !ok || user.Name == "" || user.Email == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
