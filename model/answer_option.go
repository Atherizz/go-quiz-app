package model

import "gorm.io/gorm"

type AnswerOption struct {
	gorm.Model
	QuestionId int
	Question   Question `gorm:"foreignKey:QuestionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OptionNumber int
	OptionText string
	IsCorrect  bool
}
