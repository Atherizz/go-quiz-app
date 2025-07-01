package repository

import (
	"context"
	"fmt"
	"google-oauth/model"
	"log"

	"gorm.io/gorm"
)

type AnswerOptionRepository struct {
}

func NewAnswerOptionRepository() *AnswerOptionRepository {
	return &AnswerOptionRepository{}
}

func (repo *AnswerOptionRepository) Insert(ctx context.Context, db *gorm.DB, answerOption model.AnswerOption) (model.AnswerOption, error) {
	newAnswerOption := model.AnswerOption{
		QuestionId: answerOption.QuestionId,
		OptionText: answerOption.OptionText,
		IsCorrect: answerOption.IsCorrect,
		OptionNumber: answerOption.OptionNumber,
	}

	result := db.Create(&newAnswerOption)

	if result.Error != nil {
		log.Printf("Error creating AnswerOption: %v", result.Error)
		return model.AnswerOption{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.AnswerOption{}, result.Error
	}

	return newAnswerOption, nil
}

func (repo *AnswerOptionRepository) Update(ctx context.Context, db *gorm.DB, answerOption model.AnswerOption) (model.AnswerOption, error) {

	updatedAnswerOption := model.AnswerOption{
		Model: gorm.Model{ID: answerOption.ID},
		QuestionId: answerOption.QuestionId,
		OptionText: answerOption.OptionText,
		IsCorrect: answerOption.IsCorrect,
		OptionNumber: answerOption.OptionNumber,
	}
	result := db.Model(&updatedAnswerOption).Updates(model.AnswerOption{QuestionId: updatedAnswerOption.QuestionId, OptionText: updatedAnswerOption.OptionText, IsCorrect: updatedAnswerOption.IsCorrect})

	if result.Error != nil {
		log.Printf("Error creating AnswerOption: %v", result.Error)
		return model.AnswerOption{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.AnswerOption{}, result.Error
	}

	return updatedAnswerOption, nil
}


func (repo *AnswerOptionRepository) GetAnswerOptionGroupByQuestion(ctx context.Context, db *gorm.DB, questionId int) ([]model.AnswerOption, error) {

	var answerOptions []model.AnswerOption

	result := db.Model(&model.AnswerOption{}).Where("question_id = ?", questionId).Preload("Question").Find(&answerOptions)

	if result.Error != nil {
		fmt.Println("Error saat ambil data AnswerOption:", result.Error)
		return []model.AnswerOption{}, result.Error
	}

	return answerOptions, nil
}

func (repo *AnswerOptionRepository) Delete(ctx context.Context, db *gorm.DB, id int) (error) {
	result := db.Delete(&model.AnswerOption{}, id)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return  nil
}


func (repo *AnswerOptionRepository) GetAnswerOptionById(ctx context.Context, db *gorm.DB, id int) (model.AnswerOption, error) {
	var AnswerOption model.AnswerOption

	err := db.Where("id = ?", id).
		First(&AnswerOption).Error

	if err != nil {
		return model.AnswerOption{}, err
	}

	return AnswerOption, nil
}
