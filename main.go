package main

import (
	"encoding/gob"
	"google-oauth/app"
	"google-oauth/handler"
	"google-oauth/model"
	"google-oauth/repository"
	"google-oauth/service"

	"github.com/gin-contrib/cors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()

// ...existing code...
    // User
    userRepository := repository.NewUserRepository()
    userService := service.NewUserService(*userRepository, db)
    userController := handler.NewAuthHandler(userService)

    // Subject
    subjectRepository := repository.NewSubjectRepository()
    subjectService := service.NewSubjectService(*subjectRepository, db)
    subjectController := handler.NewSubjectHandler(subjectService)

    // Quiz
    quizRepository := repository.NewQuizRepository()
    quizService := service.NewQuizService(*quizRepository, db)
    quizController := handler.NewQuizHandler(quizService)

    // Question
    questionRepository := repository.NewQuestionRepository()
    questionService := service.NewQuestionService(*questionRepository, db)
    questionController := handler.NewQuestionHandler(questionService)

    // AnswerOption
    answerOptionRepository := repository.NewAnswerOptionRepository()
    answerOptionService := service.NewAnswerOptionService(*answerOptionRepository, db)
    answerOptionController := handler.NewAnswerOptionHandler(answerOptionService)

    // UserAnswer
    userAnswerRepository := repository.NewUserAnswerRepository()
    userAnswerService := service.NewUserAnswerService(*userAnswerRepository, db)
    userAnswerController := handler.NewUserAnswerHandler(userAnswerService)

    // UserQuizResult
    userQuizResultRepository := repository.NewUserQuizResultRepository()
    userQuizResultService := service.NewUserQuizResultService(*userQuizResultRepository, db)
    userQuizResultController := handler.NewUserQuizResultHandler(userQuizResultService)

    appHandler := handler.AppHandler{
        Auth:           userController,
        Subject:        subjectController,
        Quiz:           quizController,
        Question:       questionController,
        AnswerOption:   answerOptionController,
        UserAnswer:     userAnswerController,
        UserQuizResult: userQuizResultController,
    }

	router := app.NewRouter(appHandler)
	router.Use(cors.Default())
	
	gob.Register(model.User{})

	router.Run(":8000")
}
