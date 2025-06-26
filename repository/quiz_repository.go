package repository

import (
	"context"
	"fmt"
	"google-oauth/model"
	"log"

	"gorm.io/gorm"
)

type QuizRepository struct {
}

func NewQuizRepository() *QuizRepository {
	return &QuizRepository{}
}

func (repo *QuizRepository) Insert(ctx context.Context, db *gorm.DB, quiz model.Quiz) model.Quiz {

	newQuiz := model.Quiz{
		Title: quiz.Title,
		SubjectId: quiz.SubjectId,
		Description: quiz.Description,
	}

	result := db.Create(&newQuiz)

	if result.Error != nil {
		log.Printf("Error creating Quiz: %v", result.Error)
		return model.Quiz{}
	}

	if result.RowsAffected == 0 {
		return model.Quiz{}
	}

	return newQuiz
}

func (repo *QuizRepository) Update(ctx context.Context, db *gorm.DB, quiz model.Quiz) model.Quiz {

	updatedQuiz := model.Quiz{
		Model: gorm.Model{ID: quiz.ID},
		Title: quiz.Title,
		Description: quiz.Description ,
	}

	result := db.Model(&updatedQuiz).Updates(model.Quiz{Title: updatedQuiz.Title, Description: updatedQuiz.Description})

	if result.Error != nil {
		log.Printf("Error creating Quiz: %v", result.Error)
		return model.Quiz{}
	}

	if result.RowsAffected == 0 {
		return model.Quiz{}
	}

	return updatedQuiz
}

func (repo *QuizRepository) GetAll(ctx context.Context, db *gorm.DB) []model.Quiz {

	var quizs []model.Quiz

	result := db.Find(&quizs)

	if result.Error != nil {
		fmt.Println("Error saat ambil data Quiz:", result.Error)
		return []model.Quiz{}
	}

	return quizs
}

func (repo *QuizRepository) GetQuizById(ctx context.Context, db *gorm.DB, id int) (model.Quiz, error) {
	var quiz model.Quiz

	err := db.Where("id = ?", id).
		First(&quiz).Error

	if err != nil {
		return model.Quiz{}, err
	}

	return quiz, nil
}
