package repository

import (
	"akrab-bangkit2022-api/entity"

	"gorm.io/gorm"
)

type QuizReposotory interface {
	FindQuizByLevel(level string) ([]entity.Quiz, error)
}

type quizReposotory struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) QuizReposotory {
	return &quizReposotory{db}
}

func (r *quizReposotory) FindQuizByLevel(level string) ([]entity.Quiz, error) {
	var quiz []entity.Quiz

	query := r.db.Where("level = ?", level)
	err := query.Order("id  asc").Find(&quiz).Error

	if err != nil {
		return nil, err
	}
	return quiz, nil

}