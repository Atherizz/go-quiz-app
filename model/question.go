package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	QuizId       int
	Quiz         Quiz `gorm:"foreignKey:QuizId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	QuestionText string
}
