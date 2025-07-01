package model

import "gorm.io/gorm"

type UserAnswer struct {
	gorm.Model
	UserId int
	User User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	QuizId int
	Quiz Quiz `gorm:"foreignKey:QuizId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	QuestionId int
	Question Question `gorm:"foreignKey:QuestionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SelectedOption int
	IsCorrect bool
}

type QuizAnswerSummary struct {
	Score float64
	TotalQuestions  int64
	CorrectAnswers  int64
}