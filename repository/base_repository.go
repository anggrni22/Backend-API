package repository

import (
	"akrab-bangkit2022-api/config"

	"gorm.io/gorm"
)

type baseRepository struct {
	db *gorm.DB
}

type BaseRepository interface {
	GetDB() *gorm.DB
	BeginTx()
	CommitTx()
	RollbackTx()
	GetSchema() config.Schema
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{db}
}

func (br *baseRepository) GetDB() *gorm.DB {
	return br.db
}

func (br *baseRepository) BeginTx() {
	br.db = br.GetDB().Begin()
}

func (br *baseRepository) CommitTx() {
	br.GetDB().Commit()
}

func (br *baseRepository) RollbackTx() {
	br.GetDB().Rollback()
}

func (br *baseRepository) GetSchema() config.Schema {
	return config.C.Database.Schema
}
