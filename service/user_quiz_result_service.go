package service

import (
	"context"
	"google-oauth/helper"
	"google-oauth/model"
	"google-oauth/repository"
	"google-oauth/web"

	"gorm.io/gorm"
)

type UserQuizResultService struct {
	Repository repository.UserQuizResultRepository
	DB         *gorm.DB
}

func NewUserQuizResultService(repo repository.UserQuizResultRepository, db *gorm.DB) *UserQuizResultService {
	return &UserQuizResultService{
		Repository: repo,
		DB:         db,
	}
}

func (service *UserQuizResultService) GetQuizResultGroupByQuizAndUser(ctx context.Context, db *gorm.DB, quizId int, userId int) (web.UserQuizResultResponse, error) {

	var response web.UserQuizResultResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		userQuizAnswer, err := service.Repository.GetQuizResultGroupByQuizAndUser(ctx, tx, quizId, userId)

		result := model.UserQuizResult{
			UserId:         userId,
			QuizId:         quizId,
			Score:          userQuizAnswer.Score,
			TotalQuestions: int(userQuizAnswer.TotalQuestions),
			CorrectAnswers: int(userQuizAnswer.CorrectAnswers),
		}

		response = helper.ToUserQuizResultResponse(result)

		if err != nil {
			return err
		}
		return nil
	})

	return response, err
}

func (service *UserQuizResultService) GetUserQuizResultGroupByQuiz(ctx context.Context, db *gorm.DB, id int) ([]web.UserQuizResultResponse, error) {

	var responses []web.UserQuizResultResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		userQuizResults, err := service.Repository.GetUserQuizResultGroupByQuiz(ctx, db, id)

		if err != nil {
			return err
		}

		for _, u := range userQuizResults {
			responses = append(responses, helper.ToUserQuizResultResponse(u))
		}

		return nil

	})

	return responses, err

}

func (service *UserQuizResultService) GetUserQuizResultGroupByUser(ctx context.Context, db *gorm.DB, id int) ([]web.UserQuizResultResponse, error) {

	var responses []web.UserQuizResultResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		userQuizResults, err := service.Repository.GetUserQuizResultGroupByUser(ctx, db, id)

		if err != nil {
			return err
		}

		for _, u := range userQuizResults {
			responses = append(responses, helper.ToUserQuizResultResponse(u))
		}

		return nil

	})

	return responses, err

}



func (service *UserQuizResultService) Delete(ctx context.Context, db *gorm.DB, id int) error {

	err :=service.DB.Transaction(func(tx *gorm.DB) error {

		err := service.Repository.Delete(ctx, db, id)

		if err != nil {
			return err
		}

		return nil
	})

	return err

}

