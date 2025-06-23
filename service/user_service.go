package service

import (
	"context"
	"google-oauth/helper"
	"google-oauth/model"
	"google-oauth/repository"
	"google-oauth/web"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserService struct {
	Repository repository.UserRepository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewUserService(repo repository.UserRepository, db *gorm.DB, validator *validator.Validate) *UserService {
	return &UserService{
		Repository: repo,
		DB:         db,
		Validate:   validator,
	}
}

func (service *UserService) RegisterFromGoogle(ctx context.Context, request model.User) web.UserResponse {
	var response web.UserResponse

	service.DB.Transaction(func(tx *gorm.DB) error {

		user := model.User{
			// id,google_id,name,email,picture,provider,role
			GoogleId: request.GoogleId,
			Name:     request.Name,
			Email:    request.Email,
			Picture:  request.Picture,
			Provider: request.Provider,
			Role:     request.Role,
		}

		userRegister := service.Repository.RegisterFromGoogle(ctx, tx, user)
		response = helper.ToUserResponse(userRegister)

		return nil
	})

	return response
}

func (service *UserService) RegisterDefault(ctx context.Context, request web.UserRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}
	}
	var response web.UserResponse

	service.DB.Transaction(func(tx *gorm.DB) error {

		user := model.User{
			Name:     request.Name,
			Email:    request.Email,
			Password: request.Password,
		}

		userRegister := service.Repository.RegisterDefault(ctx, tx, user)
		response = helper.ToUserResponse(userRegister)

		return nil
	})

	return response
}

func (service *UserService) GetUserByEmail(ctx context.Context, email string) web.UserResponse {
	var response web.UserResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		user, err := service.Repository.GetUserByEmail(ctx, tx, email)

		if err != nil {
			return err
		}
		response = helper.ToUserResponse(user)

		return nil
	})

	if err != nil {
		return web.UserResponse{}
	}

	return response
}
