package repository

import (
	"context"
	"google-oauth/model"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) RegisterFromGoogle(ctx context.Context, db *gorm.DB, user model.User) model.User {

	newUser := model.User{
		GoogleId: user.GoogleId,
		Name:     user.Name,
		Email:    user.Email,
		Picture:  user.Picture,
		Provider: user.Provider,
		Role:     user.Role,
		LastLoginAt: time.Now(),
	}

	result := db.Create(&newUser)

if result.Error != nil {

    log.Printf("Error creating user: %v", result.Error)
    return model.User{} // Atau kembalikan (model.User{}, result.Error)
}

if result.RowsAffected == 0 {
    return model.User{}
}

	return newUser

}

func (repo *UserRepository) RegisterDefault(ctx context.Context, db *gorm.DB, user model.User) model.User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashedString := string(hashedPassword)

	newUser := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedString,
		LastLoginAt: time.Now(),
	}

	result := db.Create(&newUser)

	if result.RowsAffected == 0 {
		return model.User{}
	}

	return newUser
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, db *gorm.DB, email string) (model.User, error) {
	var user model.User

	err := db.Select("id", "google_id", "name", "email", "picture", "provider", "role").
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
