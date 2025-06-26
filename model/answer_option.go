package model

type AnswerQuestion struct {
	ID int `gorm:"primaryKey"`
	QuestionId int 
	Question Question `gorm:"foreignKey:QuestionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OptionText string
	IsCorrect bool
}
