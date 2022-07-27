package entity

import "akrab-bangkit2022-api/config"

type UserAccess struct {
	ID           int    `json:"id"`
	IDUser       int    `json:"id_user"`
	Token        string `json:"token"`
	LastUpdate   string `json:"last_update"`
}

func (UserAccess) TableName() string {
	return config.C.Database.Schema.App + ".tb_user_access"
}
