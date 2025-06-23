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
	router.GET("/logout", middleware.OauthMiddleware(), oauthController.Logout)

	router.GET("/login", handler.LoginView)
	router.GET("/register", handler.RegisterView)
	router.GET("/home", middleware.OauthMiddleware(), handler.HomeView)
	router.GET("/profile", middleware.OauthMiddleware(), middleware.AuthMiddleware(), handler.ProfileView)

	api := router.Group("/api")
	api.GET("/user", middleware.OauthMiddleware(), middleware.AuthMiddleware(), handler.ProfileApi)
	api.POST("/register", oauthController.RegisterDefault)

	return router

}
