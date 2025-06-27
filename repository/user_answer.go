package repository

import (
	"context"
	"fmt"
	"google-oauth/model"
	"log"

	"gorm.io/gorm"
)

type UserAnswerRepository struct {
}

func NewUserAnswerRepository() *UserAnswerRepository {
	return &UserAnswerRepository{}
}

func (repo *UserAnswerRepository) Insert(ctx context.Context, db *gorm.DB, userAnswer model.UserAnswer) (model.UserAnswer, error) {
	newUserAnswer := model.UserAnswer{
		UserId:         userAnswer.UserId,
		QuizId:         userAnswer.QuizId,
		QuestionId:     userAnswer.QuestionId,
		SelectedOption: userAnswer.SelectedOption,
		IsCorrect:      userAnswer.IsCorrect,
	}

	result := db.Create(&newUserAnswer)

	if result.Error != nil {
		log.Printf("Error creating UserAnswer: %v", result.Error)
		return model.UserAnswer{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.UserAnswer{}, result.Error
	}

	return newUserAnswer, nil
}

func (repo *UserAnswerRepository) GetUserAnswerGroupByQuiz(ctx context.Context, db *gorm.DB, quizId int) ([]model.UserAnswer, error) {

	var userAnswers []model.UserAnswer

	result := db.Model(&model.UserAnswer{}).Where("quiz_id = ?", quizId).Preload("User").Preload("Quiz").Preload("Question").Find(&userAnswers)

	if result.Error != nil {
		fmt.Println("Error saat ambil data UserAnswer:", result.Error)
		return []model.UserAnswer{}, result.Error
	}

	return userAnswers, nil
}

func (repo *UserAnswerRepository) Delete(ctx context.Context, db *gorm.DB, id int) error {
	result := db.Delete(&model.UserAnswer{}, id)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (repo *UserAnswerRepository) GetUserAnswerById(ctx context.Context, db *gorm.DB, id int) (model.UserAnswer, error) {
	var userAnswer model.UserAnswer

	err := db.Where("id = ?", id).
		First(&userAnswer).Error

	if err != nil {
		return model.UserAnswer{}, err
	}

	return userAnswer, nil
}
