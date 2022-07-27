package repository

import (
	"akrab-bangkit2022-api/entity"

	"gorm.io/gorm"
)

type LevelReposotory interface {
	FindLevelByLevel(level string) (*entity.Level, error)
	FindAllLevel() ([]entity.Level, error)
	FindAllLevelByTipe(tipe string) ([]entity.Level, error)
}

type levelReposotory struct {
	db *gorm.DB
}

func NewLevelRepository(db *gorm.DB) LevelReposotory {
	return &levelReposotory{db}
}

func (r *levelReposotory) FindLevelByLevel(level string) (*entity.Level, error) {
	var findlevel entity.Level

	query := r.db.Where("level = ?", level)
	err := query.Find(&findlevel).Error

	if err != nil {
		return nil, err
	}
	return &findlevel, nil

}

func (r *levelReposotory) FindAllLevel() ([]entity.Level, error) {
	var level []entity.Level

	err := r.db.Find(&level).Error

	if err != nil {
		return nil, err
	}
	return level, nil
}

func (r *levelReposotory) FindAllLevelByTipe(tipe string) ([]entity.Level, error) {
	var findTipe []entity.Level

	query := r.db.Where("tipe = ?", tipe)
	err := query.Order("id  asc").Find(&findTipe).Error

	if err != nil {
		return nil, err
	}
	return findTipe, nil
}