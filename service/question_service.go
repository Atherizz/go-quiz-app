package service

import (
	"context"
	"google-oauth/helper"
	"google-oauth/model"
	"google-oauth/repository"
	"google-oauth/web"
	"gorm.io/gorm"
)

type QuestionService struct {
	Repository repository.QuestionRepository
	DB         *gorm.DB
}

func NewQuestionService(repo repository.QuestionRepository, db *gorm.DB) *QuestionService {
	return &QuestionService{
		Repository: repo,
		DB:         db,
	}
}

func (service *QuestionService) Insert(ctx context.Context, request web.QuestionRequest) (web.QuestionResponse, error) {

	var response web.QuestionResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		question := model.Question{
			QuizId: request.QuizId,
			QuestionText: request.QuestionText,
		}

		newQuestion, err := service.Repository.Insert(ctx, tx, question)

		if err != nil {
		return err
		}
		
		response = helper.ToQuestionResponse(newQuestion)

		return nil
	})

	return response, err

}

func (service *QuestionService) Update(ctx context.Context, request web.QuestionRequest) (web.QuestionResponse, error) {

	var response web.QuestionResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		Question := model.Question{
			Model: gorm.Model{ID: uint(request.ID)},
			QuizId: request.QuizId,
			QuestionText: request.QuestionText,
		}

		updatedQuestion, err := service.Repository.Update(ctx, tx, Question)

		if err != nil {
		return err
		}
		
		response = helper.ToQuestionResponse(updatedQuestion)

		return nil
	})

	return response, err

}

func (service *QuestionService) GetQuestionGroupByQuiz(ctx context.Context, db *gorm.DB, id int) ([]web.QuestionResponse, error) {

	
	var responses []web.QuestionResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		questions, err := service.Repository.GetQuestionGroupByQuiz(ctx, tx, id)

		if err != nil {
		return err
		}

		for _, question := range questions {
			responses = append(responses, helper.ToQuestionResponse(question))
		}
		
		return nil
	})

	return responses, err
}

func (service *QuestionService) Delete(ctx context.Context, db *gorm.DB, id int) (error) {


	err := service.DB.Transaction(func(tx *gorm.DB) error {

		err := service.Repository.Delete(ctx, db, id)

		if err != nil {
		return err
		}
		
		return nil
	})

	return err

}

func (service *QuestionService) GetQuestionById(ctx context.Context, db *gorm.DB, id int) (web.QuestionResponse, error) {

	var response web.QuestionResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		Question, err := service.Repository.GetQuestionById(ctx, db, id)

		if err != nil {
		return err
		}

		response = helper.ToQuestionResponse(Question)

		return nil

	})

	return response, err

}

