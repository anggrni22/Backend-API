package repository

import (
	"akrab-bangkit2022-api/entity"
	base "akrab-bangkit2022-api/repository"
)

type userDataRepository struct {
	base base.BaseRepository
}

type UserDataRepository interface {
	Create(user entity.UserData) (entity.UserData, error)
	Find(id int) (*entity.UserData, error)
}

func NewUserDataRepository(ar base.BaseRepository) UserDataRepository {
	return &userDataRepository{ar}
}

func (r *userDataRepository) Create(user entity.UserData) (entity.UserData, error) {
	err := r.base.GetDB().Create(&user).Error
	return user, err
}

func (r *userDataRepository) Find(id int) (*entity.UserData, error) {
	var users entity.UserData
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&entity.UserData{IDUser: id})

	err := query.First(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}
