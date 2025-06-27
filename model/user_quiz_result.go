package model

import "gorm.io/gorm"

type UserQuizResult struct {
	gorm.Model
	UserId int
	User User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	QuizId int
	Quiz Quiz `gorm:"foreignKey:QuizId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Score int
	TotalQuestions int
	CorrectAnswers int
}