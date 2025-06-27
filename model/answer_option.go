package model

import "gorm.io/gorm"

type AnswerOption struct {
	QuestionId int
	Question   Question `gorm:"foreignKey:QuestionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OptionText string
	IsCorrect  bool
	gorm.Model
}
