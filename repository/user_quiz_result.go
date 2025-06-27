package repository

import (
	"context"
	"fmt"
	"google-oauth/model"
	"log"

	"gorm.io/gorm"
)

type UserQuizResultRepository struct {
}

func NewUserQuizResultRepository() *UserQuizResultRepository {
	return &UserQuizResultRepository{}
}

func (repo *UserQuizResultRepository) Insert(ctx context.Context, db *gorm.DB, userQuizResult model.UserQuizResult) (model.UserQuizResult, error) {
	// CorrectAnswers int
	newUserQuizResult := model.UserQuizResult{
		UserId:         userQuizResult.UserId,
		QuizId:         userQuizResult.QuizId,
		Score: userQuizResult.Score,
		TotalQuestions: userQuizResult.TotalQuestions,
		CorrectAnswers: userQuizResult.CorrectAnswers,
	}

	result := db.Create(&newUserQuizResult)

	if result.Error != nil {
		log.Printf("Error creating UserQuizResult: %v", result.Error)
		return model.UserQuizResult{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.UserQuizResult{}, result.Error
	}

	return newUserQuizResult, nil
}

func (repo *UserQuizResultRepository) GetUserQuizResultGroupByUser(ctx context.Context, db *gorm.DB, userId int) ([]model.UserQuizResult, error) {

	var userQuizResults []model.UserQuizResult

	result := db.Model(&model.UserQuizResult{}).Where("user_id = ?", userId).Preload("User").Preload("Quiz").Find(&userQuizResults)

	if result.Error != nil {
		fmt.Println("Error saat ambil data UserQuizResult:", result.Error)
		return []model.UserQuizResult{}, result.Error
	}

	return userQuizResults, nil
}

func (repo *UserQuizResultRepository) Delete(ctx context.Context, db *gorm.DB, id int) error {
	result := db.Delete(&model.UserQuizResult{}, id)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (repo *UserQuizResultRepository) GetUserQuizResultById(ctx context.Context, db *gorm.DB, id int) (model.UserQuizResult, error) {
	var UserQuizResult model.UserQuizResult

	err := db.Where("id = ?", id).
		First(&UserQuizResult).Error

	if err != nil {
		return model.UserQuizResult{}, err
	}

	return UserQuizResult, nil
}
