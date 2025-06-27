package model

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	SubjectId int 
	Subject   Subject `gorm:"foreignKey:SubjectId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Questions []Question
	Description string
	Title       string
}