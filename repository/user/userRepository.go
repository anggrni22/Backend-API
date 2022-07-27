package repository

import (
	"akrab-bangkit2022-api/entity"
	base "akrab-bangkit2022-api/repository"
)

type usersRepository struct {
	base base.BaseRepository
}

type UsersRepository interface {
	Create(user entity.User) (entity.User, error)
	Find(id int) (*entity.User, error)
	FindByParam(filter map[string]interface{}) (*entity.User, error)
}

func NewUsersRepository(ar base.BaseRepository) UsersRepository{
	return &usersRepository{ar}
}

func (r *usersRepository) Create(user entity.User) (entity.User, error) {
	err := r.base.GetDB().Create(&user).Error
	return user, err
}

func (r *usersRepository) FindByParam(filter map[string]interface{}) (*entity.User, error) {
	var users entity.User
	query := r.base.GetDB()

	if filter["email"] != nil {
		query = query.Where(&entity.User{Email: filter["email"].(string)})
	}
	err := query.First(&users).Error
	if err != nil {
		return &users, err
	}

	return &users, nil
}

func (r *usersRepository) Find(id int) (*entity.User, error) {
	var users entity.User
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&entity.User{IdUser: id})

	err := query.First(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}
