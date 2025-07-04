package service

import (
	"context"
	"google-oauth/helper"
	"google-oauth/repository"
	"google-oauth/web"

	"gorm.io/gorm"
)

type UserAnswerService struct {
	Repository repository.UserAnswerRepository
	DB         *gorm.DB
}

func NewUserAnswerService(repo repository.UserAnswerRepository, db *gorm.DB) *UserAnswerService {
	return &UserAnswerService{
		Repository: repo,
		DB:         db,
	}
}

func (service *UserAnswerService) SaveAllAnswers(ctx context.Context, request web.SubmitQuizRequest, quizId int) (web.UserQuizResultResponse, error) {

	var response web.UserQuizResultResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		userAnswer := web.SubmitQuizRequest{
			UserId:  request.UserId,
			QuizId:  quizId,
			Answers: request.Answers,
		}

		quizResult, err := service.Repository.SaveAllAnswers(ctx, tx, userAnswer)
		if err != nil {
			return err
		}

		response = helper.ToUserQuizResultResponse(quizResult)

		return nil
	})

	return response, err

}

func (service *UserAnswerService) Delete(ctx context.Context, id int) error {

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		err := service.Repository.Delete(ctx, tx, id)

		if err != nil {
			return err
		}

		return nil
	})

	return err

}
