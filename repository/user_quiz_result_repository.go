package repository

import (
	"context"
	"fmt"
	"google-oauth/model"

	"gorm.io/gorm"
)

type UserQuizResultRepository struct {
}

func NewUserQuizResultRepository() *UserQuizResultRepository {
	return &UserQuizResultRepository{}
}

func (repo *UserQuizResultRepository) GetQuizResultGroupByQuizAndUser(ctx context.Context, db *gorm.DB, quizId int, userId int) (model.UserQuizResult, error) {
	var userQuizResult model.UserQuizResult

	result := db.Model(&model.UserQuizResult{}).Where("quiz_id = ? AND user_id = ?", quizId, userId).First(&userQuizResult)

	if result.Error != nil {
		fmt.Println("Error saat ambil data UserQuizResult:", result.Error)
		return model.UserQuizResult{}, result.Error
	}

	var userResult model.UserQuizResult

	err := db.
	Preload("User").
	Preload("Quiz").
	First(&userResult, userQuizResult.ID).Error

	if err != nil {
		fmt.Println("Error saat ambil data UserResult:", result.Error)
		return model.UserQuizResult{}, result.Error
	}

	return userResult, nil
}

func (repo *UserQuizResultRepository) GetUserQuizResultGroupByQuiz(ctx context.Context, db *gorm.DB, quizId int) ([]model.UserQuizResult, error) {

	var userQuizResults []model.UserQuizResult

	result := db.Model(&model.UserQuizResult{}).Where("quiz_id = ?", quizId).Preload("User").Preload("Quiz").Find(&userQuizResults)

	if result.Error != nil {
		fmt.Println("Error saat ambil data UserQuizResult:", result.Error)
		return []model.UserQuizResult{}, result.Error
	}

	return userQuizResults, nil
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
