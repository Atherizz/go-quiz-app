package service

import (
	"context"
	"google-oauth/helper"
	"google-oauth/model"
	"google-oauth/repository"
	"google-oauth/web"
	"gorm.io/gorm"
)

type SubjectService struct {
	Repository repository.SubjectRepository
	DB         *gorm.DB
}

func NewSubjectService(repo repository.SubjectRepository, db *gorm.DB) *SubjectService {
	return &SubjectService{
		Repository: repo,
		DB:         db,
	}
}

func (service *SubjectService) Insert(ctx context.Context, request web.SubjectRequest) (web.SubjectResponse, error) {
	
	var response web.SubjectResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		subject := model.Subject{
			SubjectName: request.SubjectName,
		}

		newSubject, err := service.Repository.Insert(ctx, tx, subject)

		if err != nil {
		return err
		}
		
		response = helper.ToSubjectResponse(newSubject)

		return nil
	})

	return response, err

}

func (service *SubjectService) Update(ctx context.Context, request web.SubjectRequest) (web.SubjectResponse, error) {
	
	var response web.SubjectResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		subject := model.Subject{
			Model: gorm.Model{ID: uint(request.ID)},
			SubjectName: request.SubjectName,
		}

		updatedSubject, err := service.Repository.Update(ctx, tx, subject)

		if err != nil {
		return err
		}
		
		response = helper.ToSubjectResponse(updatedSubject)

		return nil
	})

	return response, err

}

func (service *SubjectService) GetAll(ctx context.Context, db *gorm.DB) ([]web.SubjectResponse, error) {

	
	var responses []web.SubjectResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		subjects, err := service.Repository.GetAll(ctx, tx)

		if err != nil {
		return err
		}

		for _, subject := range subjects {
			responses = append(responses, helper.ToSubjectResponse(subject))
		}
		
		return nil
	})

	return responses, err
}

func (service *SubjectService) Delete(ctx context.Context, db *gorm.DB, id int) (error) {

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		err := service.Repository.Delete(ctx, db, id)

		if err != nil {
		return err
		}
		
		return nil
	})

	return err

}

func (service *SubjectService) GetSubjectById(ctx context.Context, db *gorm.DB, id int) (web.SubjectResponse, error) {

	var response web.SubjectResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		subject, err := service.Repository.GetSubjectById(ctx, db, id)

		if err != nil {
		return err
		}

		response = helper.ToSubjectResponse(subject)

		return nil

	})

	return response, err

}

