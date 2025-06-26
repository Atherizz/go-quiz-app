package app

import (
	"google-oauth/handler"
	"google-oauth/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(oauthController *handler.OauthController) *gin.Engine {
	router := gin.Default()

	router.GET("/auth/google/login", oauthController.LoginOauth)
	router.GET("/callback", oauthController.Callback)
	router.GET("/login", handler.LoginView)
	router.GET("/register", handler.RegisterView)

	api := router.Group("/api")
	{
		api.GET("/user", middleware.OauthMiddleware(), middleware.AuthMiddleware(), handler.ProfileApi)
		api.POST("/register", oauthController.RegisterDefault)
	}

	oauthGroup := router.Group("/")
	oauthGroup.Use(middleware.OauthMiddleware())
	{
		oauthGroup.GET("/home", handler.HomeView)
		oauthGroup.GET("/logout", oauthController.Logout)

		secured := oauthGroup.Group("/")
		secured.Use(middleware.AuthMiddleware())
		{
			secured.GET("/profile", handler.ProfileView)
		}
	}

	return router
}
