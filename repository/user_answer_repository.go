package repository

import (
	"context"
	"google-oauth/model"
	"google-oauth/web"
	"log"
	"math"

	"gorm.io/gorm"
)

type UserAnswerRepository struct {
}

func NewUserAnswerRepository() *UserAnswerRepository {
	return &UserAnswerRepository{}
}


func (repo *UserAnswerRepository) SaveAllAnswers(ctx context.Context, db *gorm.DB, userAnswer web.SubmitQuizRequest) (model.UserQuizResult, error) {

	var bulkInsertData []map[string]interface{}
	var correctAnswer model.AnswerOption

	for _, a := range userAnswer.Answers {
		db.Where(&model.AnswerOption{QuestionId: a.QuestionId, IsCorrect: true}).First(&correctAnswer)

		isCorrect := correctAnswer.OptionNumber == a.SelectedOption

		bulkInsertData = append(bulkInsertData, map[string]interface{}{
			"UserId":         userAnswer.UserId,
			"QuizId":         userAnswer.QuizId,
			"QuestionId":     a.QuestionId,
			"SelectedOption": a.SelectedOption,
			"IsCorrect":      isCorrect,
		})
	}

	result := db.Model(&model.UserAnswer{}).Create(bulkInsertData)

	if result.Error != nil {
		log.Printf("Error creating UserAnswer: %v", result.Error)
		return model.UserQuizResult{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.UserQuizResult{}, result.Error
	}

	var score float64
	var totalQuestions int64
	var correctAnswers int64

	err1 := db.Model(&model.Question{}).
		Where("quiz_id = ?", userAnswer.QuizId).
		Count(&totalQuestions).Error

	if err1 != nil {
		return model.UserQuizResult{}, err1
	}

	err2 := db.Model(&model.UserAnswer{}).
		Where("quiz_id = ? AND user_id = ? AND is_correct = ?", userAnswer.QuizId, userAnswer.UserId, true).
		Count(&correctAnswers).Error
	if err2 != nil {
		return model.UserQuizResult{}, err2
	}

	if totalQuestions > 0 {
	rawScore := (float64(correctAnswers) / float64(totalQuestions)) * 100
	score = math.Round(rawScore*100) / 100 
	}

	newQuizResult := model.UserQuizResult{
		UserId: userAnswer.UserId,
		QuizId: userAnswer.QuizId,
		Score: score,
		TotalQuestions: int(totalQuestions),
		CorrectAnswers: int(correctAnswers),
	}

		insertResult := db.Create(&newQuizResult)

	if insertResult.Error != nil {
		log.Printf("Error creating UserQuizResult: %v", insertResult.Error)
		return model.UserQuizResult{}, insertResult.Error
	}

	if insertResult.RowsAffected == 0 {
		return model.UserQuizResult{}, insertResult.Error
	}

	return newQuizResult, nil
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
