package repository

import (
	"context"
	"fmt"
	"google-oauth/model"
	"log"

	"gorm.io/gorm"
)

type QuestionRepository struct {
}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{}
}

func (repo *QuestionRepository) Insert(ctx context.Context, db *gorm.DB, question model.Question) model.Question {

	newQuestion := model.Question{
		QuizId: question.QuizId,
		QuestionText: question.QuestionText,
	}

	result := db.Create(&newQuestion)

	if result.Error != nil {
		log.Printf("Error creating Question: %v", result.Error)
		return model.Question{}
	}

	if result.RowsAffected == 0 {
		return model.Question{}
	}

	return newQuestion
}

func (repo *QuestionRepository) Update(ctx context.Context, db *gorm.DB, question model.Question) model.Question {

	updatedQuestion := model.Question{
		Model: gorm.Model{ID: question.ID},
		QuestionText: question.QuestionText,
	}
	result := db.Model(&updatedQuestion).Where("id", updatedQuestion.ID).Update("question_text", question.QuestionText)

	if result.Error != nil {
		log.Printf("Error creating Question: %v", result.Error)
		return model.Question{}
	}

	if result.RowsAffected == 0 {
		return model.Question{}
	}

	return updatedQuestion
}

func (repo *QuestionRepository) GetAll(ctx context.Context, db *gorm.DB) []model.Question {

	var questions []model.Question

	result := db.Find(&questions)

	if result.Error != nil {
		fmt.Println("Error saat ambil data Question:", result.Error)
		return []model.Question{}
	}

	return questions
}

func (repo *QuestionRepository) GetQuestionById(ctx context.Context, db *gorm.DB, id int) (model.Question, error) {
	var question model.Question

	err := db.Where("id = ?", id).
		First(&question).Error

	if err != nil {
		return model.Question{}, err
	}

	return question, nil
}
