package main

import (
	"context"
	"encoding/gob"
	"google-oauth/app"
	"google-oauth/handler"
	"google-oauth/helper"
	"google-oauth/model"
	"google-oauth/repository"
	"google-oauth/service"

	"github.com/gin-contrib/cors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

    ctx := context.Background()
    client := helper.Client

	db := app.NewDB()
    
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(*userRepository, db)
	userController := handler.NewAuthHandler(userService)
    
	subjectRepository := repository.NewSubjectRepository()
	subjectService := service.NewSubjectService(*subjectRepository, db)
	subjectController := handler.NewSubjectHandler(subjectService)
    
	quizRepository := repository.NewQuizRepository()
	quizService := service.NewQuizService(*quizRepository, db)
	quizController := handler.NewQuizHandler(quizService)
    
	questionRepository := repository.NewQuestionRepository()
	questionService := service.NewQuestionService(*questionRepository, db)
	questionController := handler.NewQuestionHandler(questionService)
    
	answerOptionRepository := repository.NewAnswerOptionRepository()
	answerOptionService := service.NewAnswerOptionService(*answerOptionRepository, db)
	answerOptionController := handler.NewAnswerOptionHandler(answerOptionService)
    
	userAnswerRepository := repository.NewUserAnswerRepository()
	userAnswerService := service.NewUserAnswerService(*userAnswerRepository, db)
	userAnswerController := handler.NewUserAnswerHandler(userAnswerService)
    
	userQuizResultRepository := repository.NewUserQuizResultRepository()
	userQuizResultService := service.NewUserQuizResultService(*userQuizResultRepository, db)
	userQuizResultController := handler.NewUserQuizResultHandler(userQuizResultService)
    
    _ = client.XGroupCreateMkStream(ctx, "quiz_answer_stream", "group-1", "0")
    go app.ConsumeAnswers(ctx, userAnswerService)

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
