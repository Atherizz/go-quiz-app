package repository

import (
	"context"
	"fmt"
	"google-oauth/model"
	"log"

	"gorm.io/gorm"
)

type SubjectRepository struct {
}

func NewSubjectRepository() *SubjectRepository {
	return &SubjectRepository{}
}

func (repo *SubjectRepository) Insert(ctx context.Context, db *gorm.DB, subject model.Subject) (model.Subject, error) {

	newSubject := model.Subject{
		SubjectName: subject.SubjectName,
	}

	result := db.Create(&newSubject)

	if result.Error != nil {
		log.Printf("Error creating subject: %v", result.Error)
		return model.Subject{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Subject{}, result.Error
	}

	return newSubject, nil
}

func (repo *SubjectRepository) Update(ctx context.Context, db *gorm.DB, subject model.Subject) (model.Subject, error) {

	updatedSubject := model.Subject{
		Model:       gorm.Model{ID: subject.ID},
		SubjectName: subject.SubjectName,
	}

	result := db.Model(&updatedSubject).Where("id", updatedSubject.ID).Update("subject_name", subject.SubjectName)

	if result.Error != nil {
		log.Printf("Error creating subject: %v", result.Error)
		return model.Subject{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Subject{}, result.Error
	}

	return updatedSubject, nil
}

func (repo *SubjectRepository) GetAll(ctx context.Context, db *gorm.DB) ([]model.Subject, error) {

	var subjects []model.Subject

	result := db.Model(&model.Subject{}).Preload("Quiz").Find(&subjects)

	if result.Error != nil {
		fmt.Println("Error saat ambil data subject:", result.Error)
		return []model.Subject{}, result.Error
	}

	return subjects, nil
}

func (repo *SubjectRepository) Delete(ctx context.Context, db *gorm.DB, id int) (error) {
	result := db.Delete(&model.Subject{}, id)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return  nil
}

func (repo *SubjectRepository) GetSubjectById(ctx context.Context, db *gorm.DB, id int) (model.Subject, error) {
	var subject model.Subject

	err := db.Where("id = ?", id).
		First(&subject).Error

	if err != nil {
		return model.Subject{}, err
	}

	return subject, nil
}
