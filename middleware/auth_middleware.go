package middleware

import (
	"google-oauth/helper"
	"google-oauth/model"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := helper.Store.Get(c.Request, "user_info")
		user, ok := session.Values["user"].(model.User)

		if !ok || user.Email == "" {
			// Cek apakah request ke endpoint API
			if strings.HasPrefix(c.FullPath(), "/api") {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			} else {
				c.Redirect(http.StatusSeeOther, "/login")
			}
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

