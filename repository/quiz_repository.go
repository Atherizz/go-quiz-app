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

func (repo *QuizRepository) Insert(ctx context.Context, db *gorm.DB, quiz model.Quiz) (model.Quiz, error) {

	newQuiz := model.Quiz{
		Title:       quiz.Title,
		SubjectId:   quiz.SubjectId,
		Description: quiz.Description,
	}

	result := db.Create(&newQuiz)

	if result.Error != nil {
		log.Printf("Error creating Quiz: %v", result.Error)
		return model.Quiz{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Quiz{}, result.Error
	}

	return newQuiz, nil
}

func (repo *QuizRepository) Update(ctx context.Context, db *gorm.DB, quiz model.Quiz) (model.Quiz, error) {

	updatedQuiz := model.Quiz{
		Model:       gorm.Model{ID: quiz.ID},
		Title:       quiz.Title,
		Description: quiz.Description,
	}

	result := db.Model(&updatedQuiz).Updates(model.Quiz{Title: updatedQuiz.Title, Description: updatedQuiz.Description})

	if result.Error != nil {
		log.Printf("Error creating Quiz: %v", result.Error)
		return model.Quiz{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Quiz{}, result.Error
	}

	return updatedQuiz, nil
}

func (repo *QuizRepository) GetAll(ctx context.Context, db *gorm.DB) ([]model.Quiz, error) {

	var quizzes []model.Quiz

	result := db.Model(&model.Quiz{}).Preload("Questions").Find(&quizzes)

	if result.Error != nil {
		fmt.Println("Error saat ambil data Quiz:", result.Error)
		return []model.Quiz{}, result.Error
	}

	return quizzes, nil
}

func (repo *QuizRepository) GetQuizGroupBySubject(ctx context.Context, db *gorm.DB, idSubject int) ([]model.Quiz, error) {

	var quizzes []model.Quiz

	result := db.Model(&model.Quiz{}).Where("subject_id = ?", idSubject).Preload("Questions").Preload("Subject").Find(&quizzes)

	if result.Error != nil {
		fmt.Println("Error saat ambil data Quiz:", result.Error)
		return []model.Quiz{}, result.Error
	}

	return quizzes, nil
}

func (repo *QuizRepository) Delete(ctx context.Context, db *gorm.DB, id int) error {
	result := db.Delete(&model.Quiz{}, id)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
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
