package app

import (
	"google-oauth/handler"
	"google-oauth/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(appHandler handler.AppHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/auth/google/login", appHandler.Auth.LoginOauth)
	router.GET("/callback", appHandler.Auth.Callback)
	router.GET("/login", handler.LoginView)
	router.GET("/register", handler.RegisterView)

	api := router.Group("/api")
	{
		api.POST("/register", appHandler.Auth.RegisterDefault)

		protected := api.Group("/")
		// protected.Use(middleware.OauthMiddleware(), middleware.AuthMiddleware())
		{
			protected.GET("/user", handler.GetUserProfile)

			protected.GET("/subject", appHandler.Subject.GetAll)
			protected.GET("/subject/:id", appHandler.Subject.GetSubjectById)
			protected.DELETE("/subject/:id", appHandler.Subject.Delete)
			protected.POST("/subject", appHandler.Subject.Insert)
			protected.PUT("/subject/:id", appHandler.Subject.Update)

			protected.GET("/subjects/:subjectId/quizzes", appHandler.Quiz.GetQuizGroupBySubject)
			protected.POST("/subjects/:subjectId/quizzes", appHandler.Quiz.Insert)
			protected.GET("/quizzes/:id", appHandler.Quiz.GetQuizById)
			protected.PUT("/quizzes/:id", appHandler.Quiz.Update)
			protected.DELETE("/quizzes/:id", appHandler.Quiz.Delete)

			protected.GET("/quizzes/:quizId/questions", appHandler.Question.GetQuestionGroupByQuiz)
			protected.POST("/quizzes/:quizId/questions", appHandler.Question.Insert)
			protected.GET("/questions/:id", appHandler.Question.GetQuestionById)
			protected.PUT("/questions/:id", appHandler.Question.Update)
			protected.DELETE("/questions/:id", appHandler.Question.Delete)

	

		}
	}

	oauthGroup := router.Group("/")
	oauthGroup.Use(middleware.OauthMiddleware())

	oauthGroup.GET("/home", handler.HomeView)
	oauthGroup.GET("/logout", appHandler.Auth.Logout)

	securedGroup := oauthGroup.Group("/")
	securedGroup.Use(middleware.AuthMiddleware())

	securedGroup.GET("/profile", handler.ProfileView)

	return router
}
