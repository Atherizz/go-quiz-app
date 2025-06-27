package app

import (
	"google-oauth/helper"
	"google-oauth/model"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dbName := helper.LoadEnv("DB_NAME")
	port := helper.LoadEnv("PORT")
	dbUser := helper.LoadEnv("DB_USER")
	dsn := dbUser + "@tcp(localhost:" + port + ")/" + dbName + "?parseTime=true&loc=Asia%2FJakarta"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	err = db.AutoMigrate(&model.User{}, &model.Subject{}, &model.Quiz{}, &model.Question{}, &model.AnswerOption{}, &model.UserAnswer{}, &model.UserQuizResult{})
	if err != nil {
		log.Fatal("Gagal migrasi:", err)
	}

	sql, err := db.DB()
	if err != nil {
		panic(err)
	}

	sql.SetMaxIdleConns(5)
	sql.SetMaxOpenConns(20)
	sql.SetConnMaxLifetime(60 * time.Minute)
	sql.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
