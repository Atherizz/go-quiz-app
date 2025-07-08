package middleware

import (
	"encoding/json"
	"google-oauth/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		sessionId, err := c.Cookie("session_id")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		client := helper.Client

		result, err := client.Get(c.Request.Context(), "session:"+sessionId).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		userSession := helper.UserSession{}

		err = json.Unmarshal([]byte(result), &userSession)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			return
		}

		if userSession.Email == "" {
			// Cek apakah request ke endpoint API
			if strings.HasPrefix(c.FullPath(), "/api") {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			} else {
				c.Redirect(http.StatusSeeOther, "/login")
			}
			return
		}

		c.Set("user", userSession)
		c.Next()
	}
}
