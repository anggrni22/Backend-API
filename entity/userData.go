package entity

import "akrab-bangkit2022-api/config"

type UserData struct {
	IDUser       int    `json:"id_user"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	
}

func (UserData) TableName() string {
	return config.C.Database.Schema.App + ".tb_user_data"
}
