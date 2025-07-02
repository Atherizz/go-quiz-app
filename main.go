package main

import (
	"encoding/gob"
	"google-oauth/app"
	"google-oauth/handler"
	"google-oauth/model"
	"google-oauth/repository"
	"google-oauth/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()


	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(*userRepository, db)
	userController := handler.NewAuthHandler(userService)

	subjectRepository := repository.NewSubjectRepository()
	subjectService := service.NewSubjectService(*subjectRepository, db)
	subjectController := handler.NewSubjectHandler(subjectService)

	appHandler := handler.AppHandler{
		Auth: userController,
		Subject: subjectController,

	}

	router := app.NewRouter(appHandler)
	gob.Register(model.User{})

	router.Run(":8000")
}
