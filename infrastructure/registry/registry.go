package registry

import (
	"akrab-bangkit2022-api/controller"

	"gorm.io/gorm"
)

type registry struct {
	db  *gorm.DB
}

type Registry interface {
	NewModulAndQuizController() controller.ModulAndQuizController
	NewUsersController() controller.UserController	
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}
