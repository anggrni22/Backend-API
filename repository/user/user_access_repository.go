package repository

import (
	"time"

	user "akrab-bangkit2022-api/entity"
	base "akrab-bangkit2022-api/repository"
)

type userAccessRepository struct {
	base base.BaseRepository
}

type UserAccessRepository interface {
	Find(id int) (user.UserAccess, error)
	Create(ud user.UserAccess) error
	Update(id int, token string) error
	ValidToken(token string) bool
	// ValidTokenWithID(iduser int, token string) bool
	Delete(id int) error
}

func NewUserAccessRepository(ar base.BaseRepository) UserAccessRepository {
	return &userAccessRepository{ar}
}
func (r *userAccessRepository)Delete(id int) error{
	var Delete user.UserAccess
	err := r.base.GetDB().
	Where("id_user = ?", id).
	Delete(&Delete).Error
	if err != nil{
		return err
	}
	return nil
}

func (r *userAccessRepository) Find(iduser int) (user.UserAccess, error) {
	var ua user.UserAccess
	err := r.base.GetDB().Where("id_user = ?", iduser).
		First(&ua).Error
	if err != nil {
		return ua, err
	}
	return ua, nil
}
// func (r *userAccessRepository) ValidTokenWithID(iduser int, token string) bool {
// 	var ua user.UserAccess
// 	sql := r.base.GetDB().Where("id = ? and token = ?", iduser, token).
// 		First(&ua)
// 	if sql.Error != nil {
// 		return false
// 	}
// 	if sql.RowsAffected == 0 {
// 		return false
// 	}
// 	return true
// }
func (r *userAccessRepository) ValidToken(token string) bool {
	var ua user.UserAccess
	sql := r.base.GetDB().Where("token = ?", token).
		First(&ua)
	if sql.Error != nil {
		return false
	}
	if sql.RowsAffected == 0 {
		return false
	}
	return true
}
func (r *userAccessRepository) Create(ua user.UserAccess) error {
	ua.LastUpdate = time.Now().Format("2017.09.07 17:06:06")
	err := r.base.GetDB().Create(&ua).Error
	return err
}
func (r *userAccessRepository) Update(iduser int, token string) error {
	err := r.base.GetDB().Model(user.UserAccess{}).
		Where(user.UserAccess{IDUser: iduser}).
		Updates(user.UserAccess{Token: token}).Error
	if err != nil {
		return err
	}
	return nil
}
