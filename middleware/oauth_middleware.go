package middleware

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var OauthConfig = oauth2.Config{
	ClientID:     "1035471242348-e1n7ujn46982ibko0s4v3mhf54lbt12n.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "http://localhost:8000/callback",
	Scopes:       []string{"openid", "profile", "email"},
	Endpoint:     google.Endpoint,
}


func loadTokenFromRequest(c *gin.Context) (*oauth2.Token, error) {
	// get data from cookies
	cookie, err := c.Cookie("oauth_token")
	if err != nil {
		return nil, err
	}

	// Mengambil string yang di-encode base64, lalu decode ke bytes.
	tokenBytes, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		return nil, err
	}

	var token oauth2.Token
	// Mengambil JSON (byte format) lalu mengubahnya ke struct.
	err = json.Unmarshal(tokenBytes, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func OauthMiddleware() gin.HandlerFunc {
  	return func(c *gin.Context) {
		token, err := loadTokenFromRequest(c)
		if err != nil || !token.Valid() {
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}

		tokenSource := OauthConfig.TokenSource(c.Request.Context(), token)

		token, err = tokenSource.Token()
		if err != nil || !token.Valid() {
			c.Redirect(http.StatusNotFound, OauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline))
			return
		}

		c.Next()

	}
}
