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

func (repo *QuestionRepository) Insert(ctx context.Context, db *gorm.DB, question model.Question) (model.Question,error) {

	newQuestion := model.Question{
		QuizId: question.QuizId,
		QuestionText: question.QuestionText,
	}

	result := db.Create(&newQuestion)

	if result.Error != nil {
		log.Printf("Error creating Question: %v", result.Error)
		return model.Question{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Question{}, result.Error
	}

	return newQuestion, nil
}

func (repo *QuestionRepository) Update(ctx context.Context, db *gorm.DB, question model.Question) (model.Question,error) {

	updatedQuestion := model.Question{
		Model: gorm.Model{ID: question.ID},
		QuestionText: question.QuestionText,
	}
	result := db.Model(&updatedQuestion).Where("id", updatedQuestion.ID).Update("question_text", question.QuestionText)

	if result.Error != nil {
		log.Printf("Error creating Question: %v", result.Error)
		return model.Question{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Question{}, result.Error
	}

	return updatedQuestion, nil
}

func (repo *QuestionRepository) GetQuestionGroupByQuiz(ctx context.Context, db *gorm.DB, quizId int) ([]model.Question,error) {

	var questions []model.Question

	result := db.Model(&model.Question{}).Where("quiz_id = ?", quizId).Preload("AnswerOptions").Preload("Quiz").Find(&questions)

	if result.Error != nil {
		fmt.Println("Error saat ambil data Question:", result.Error)
		return []model.Question{}, result.Error
	}

	return questions, nil
}

func (repo *QuestionRepository) Delete(ctx context.Context, db *gorm.DB, id int) (error) {
	result := db.Delete(&model.Question{}, id)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return  nil
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
