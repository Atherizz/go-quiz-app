package service

import (
	"context"
	"google-oauth/helper"
	"google-oauth/model"
	"google-oauth/repository"
	"google-oauth/web"
	"gorm.io/gorm"
)

type AnswerOptionService struct {
	Repository repository.AnswerOptionRepository
	DB         *gorm.DB
}

func NewAnswerOptionService(repo repository.AnswerOptionRepository, db *gorm.DB) *AnswerOptionService {
	return &AnswerOptionService{
		Repository: repo,
		DB:         db,
	}
}

func (service *AnswerOptionService) Insert(ctx context.Context, request web.AnswerOptionRequest) (web.AnswerOptionResponse, error) {

	var response web.AnswerOptionResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		answerOption := model.AnswerOption{
			QuestionId: request.QuestionId,
			OptionText: request.OptionText,
			IsCorrect: request.IsCorrect,
		}

		newAnswerOption, err := service.Repository.Insert(ctx, tx, answerOption)

		if err != nil {
		return err
		}
		
		response = helper.ToAnswerOptionResponse(newAnswerOption)

		return nil
	})

	return response, err

}

func (service *AnswerOptionService) Update(ctx context.Context, request web.AnswerOptionRequest) (web.AnswerOptionResponse, error) {

	var response web.AnswerOptionResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		AnswerOption := model.AnswerOption{
			Model: gorm.Model{ID: uint(request.ID)},
			QuestionId: request.QuestionId,
			OptionText: request.OptionText,
			IsCorrect: request.IsCorrect,
			OptionNumber: request.OptionNumber,
		}

		updatedAnswerOption, err := service.Repository.Update(ctx, tx, AnswerOption)

		if err != nil {
		return err
		}
		
		response = helper.ToAnswerOptionResponse(updatedAnswerOption)

		return nil
	})

	return response, err

}

func (service *AnswerOptionService) GetAnswerOptionGroupByQuestion(ctx context.Context, id int) ([]web.AnswerOptionResponse, error) {

	
	var responses []web.AnswerOptionResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		AnswerOptions, err := service.Repository.GetAnswerOptionGroupByQuestion(ctx, tx, id)

		if err != nil {
		return err
		}

		for _, AnswerOption := range AnswerOptions {
			responses = append(responses, helper.ToAnswerOptionResponse(AnswerOption))
		}
		
		return nil
	})

	return responses, err
}

func (service *AnswerOptionService) Delete(ctx context.Context, id int) (error) {


	err := service.DB.Transaction(func(tx *gorm.DB) error {

		err := service.Repository.Delete(ctx, tx, id)

		if err != nil {
		return err
		}
		
		return nil
	})

	return err

}

func (service *AnswerOptionService) GetAnswerOptionById(ctx context.Context, id int) (web.AnswerOptionResponse, error) {

	var response web.AnswerOptionResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		AnswerOption, err := service.Repository.GetAnswerOptionById(ctx, tx, id)

		if err != nil {
		return err
		}

		response = helper.ToAnswerOptionResponse(AnswerOption)

		return nil

	})

	return response, err

}

