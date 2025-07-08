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

			protected.GET("/subjects", appHandler.Subject.GetAll)
			protected.GET("/subjects/:subjectId/detail", appHandler.Subject.GetSubjectById)
			protected.POST("/subjects", appHandler.Subject.Insert)
			protected.PUT("/subjects/:subjectId", appHandler.Subject.Update)
			protected.DELETE("/subjects/:subjectIid", appHandler.Subject.Delete)

			protected.GET("/subjects/:subjectId/quizzes", appHandler.Quiz.GetQuizGroupBySubject)
			protected.POST("/subjects/:subjectId/quizzes", appHandler.Quiz.Insert)
			protected.GET("/quizzes/:quizId", appHandler.Quiz.GetQuizById)
			protected.PUT("/quizzes/:quizId", appHandler.Quiz.Update)
			protected.DELETE("/quizzes/:quizId", appHandler.Quiz.Delete)

			protected.GET("/quizzes/:quizId/questions", appHandler.Question.GetQuestionGroupByQuiz)
			protected.POST("/quizzes/:quizId/questions", appHandler.Question.Insert)
			protected.GET("/questions/:questionId", appHandler.Question.GetQuestionById)
			protected.PUT("/questions/:questionId", appHandler.Question.Update)
			protected.DELETE("/questions/:questionId", appHandler.Question.Delete)

			protected.GET("/questions/:questionId/answer_options", appHandler.AnswerOption.GetAnswerOptionGroupByQuestion)
			protected.POST("/questions/:questionId/answer_options", appHandler.AnswerOption.Insert)
			protected.PUT("/answer_options/:answerOptionId", appHandler.AnswerOption.Update)
			protected.GET("/answer_options/:answerOptionId", appHandler.AnswerOption.GetAnswerOptionById)
			protected.DELETE("/answer_options/:answerOptionId", appHandler.AnswerOption.Delete)

			protected.POST("/quizzes/:quizId/user_answers", appHandler.UserAnswer.SaveAllAnswers)
			protected.DELETE("/user_answers/:userAnswerId", appHandler.UserAnswer.Delete)

			// leaderbord
			protected.GET("/quizzes/:quizId/leaderboard", appHandler.UserQuizResult.Leaderboard)
			protected.GET("/quizzes/:quizId/my_quiz_result", appHandler.UserQuizResult.GetQuizResultGroupByQuizAndUser)
			protected.GET("/my_quiz_result", appHandler.UserQuizResult.GetUserQuizResultGroupByUser)

		}
	}

	oauthGroup := router.Group("/")
	oauthGroup.Use(middleware.OauthMiddleware())

	oauthGroup.GET("/logout", appHandler.Auth.Logout)
	
	securedGroup := oauthGroup.Group("/")
	securedGroup.Use(middleware.AuthMiddleware())
	
	securedGroup.GET("/home", handler.HomeView)
	securedGroup.GET("/profile", handler.ProfileView)

	return router
}
