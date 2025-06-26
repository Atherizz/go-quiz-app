package model

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	SubjectName string
	Quiz        []Quiz
}
