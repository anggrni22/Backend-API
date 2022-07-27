package repository

import (
	"akrab-bangkit2022-api/entity"

	"gorm.io/gorm"
)

type ModulReposotory interface {
	FindModulByLevel(level string) ([]entity.Modul, error)
	FindAllModul() ([]entity.Modul, error)
	FindAllModulByTipe(tipe string) ([]entity.Modul, error)
}

type modulReposotory struct {
	db *gorm.DB
}

func NewModulRepository(db *gorm.DB) ModulReposotory {
	return &modulReposotory{db}
}

func (r *modulReposotory) FindModulByLevel(level string) ([]entity.Modul, error) {
	var modul []entity.Modul

	query := r.db.Where("level = ?", level)
	err := query.Order("id  asc").Find(&modul).Error

	if err != nil {
		return nil, err
	}
	return modul, nil

}

func (r *modulReposotory) FindAllModul() ([]entity.Modul, error) {
	var modul []entity.Modul

	err := r.db.Order("id  asc").Find(&modul).Error

	if err != nil {
		return nil, err
	}
	return modul, nil
}

func (r *modulReposotory) FindAllModulByTipe(tipe string) ([]entity.Modul, error) {
	var findTipe []entity.Modul

	query := r.db.Where("tipe = ?", tipe)
	err := query.Order("id  asc").Find(&findTipe).Error

	if err != nil {
		return nil, err
	}
	return findTipe, nil
}