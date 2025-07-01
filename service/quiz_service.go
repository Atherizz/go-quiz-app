package service

import (
	"context"
	"google-oauth/helper"
	"google-oauth/model"
	"google-oauth/repository"
	"google-oauth/web"
	"gorm.io/gorm"
)

type QuizService struct {
	Repository repository.QuizRepository
	DB         *gorm.DB
}

func NewQuizService(repo repository.QuizRepository, db *gorm.DB) *QuizService {
	return &QuizService{
		Repository: repo,
		DB:         db,
	}
}

func (service *QuizService) Insert(ctx context.Context, request web.QuizRequest) (web.QuizResponse, error) {

	var response web.QuizResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		quiz := model.Quiz{
			SubjectId: request.SubjectId,
			Description: request.Description,
			Title: request.Title,
		}

		newQuiz, err := service.Repository.Insert(ctx, tx, quiz)

		if err != nil {
		return err
		}
		
		response = helper.ToQuizResponse(newQuiz)

		return nil
	})

	return response, err

}

func (service *QuizService) Update(ctx context.Context, request web.QuizRequest) (web.QuizResponse, error) {

	var response web.QuizResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		quiz := model.Quiz{
			Model: gorm.Model{ID: uint(request.ID)},
			SubjectId: request.SubjectId,
			Description: request.Description,
			Title: request.Title,
		}

		updatedQuiz, err := service.Repository.Update(ctx, tx, quiz)

		if err != nil {
		return err
		}
		
		response = helper.ToQuizResponse(updatedQuiz)

		return nil
	})

	return response, err

}

func (service *QuizService) GetAll(ctx context.Context, db *gorm.DB) ([]web.QuizResponse, error) {

	
	var responses []web.QuizResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		quizs, err := service.Repository.GetAll(ctx, tx)

		if err != nil {
		return err
		}

		for _, quiz := range quizs {
			responses = append(responses, helper.ToQuizResponse(quiz))
		}
		
		return nil
	})

	return responses, err
}

func (service *QuizService) Delete(ctx context.Context, db *gorm.DB, id int) (error) {


	err := service.DB.Transaction(func(tx *gorm.DB) error {

		err := service.Repository.Delete(ctx, db, id)

		if err != nil {
		return err
		}
		
		return nil
	})

	return err

}

func (service *QuizService) GetQuizById(ctx context.Context, db *gorm.DB, id int) (web.QuizResponse, error) {

	var response web.QuizResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		quiz, err := service.Repository.GetQuizById(ctx, db, id)

		if err != nil {
		return err
		}

		response = helper.ToQuizResponse(quiz)

		return nil

	})

	return response, err

}

