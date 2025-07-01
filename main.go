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
	userController := handler.NewOauthController(userService)

	router := app.NewRouter(userController)
	gob.Register(model.User{})

	router.Run(":8000")
}
