package entity

import "akrab-bangkit2022-api/config"

type User struct {
	IdUser     		 int    `gorm:"primary_key:auto_increment;not_null" json:"id_user"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Salt         string `json:"salt"`
}

func (User) TableName() string {
	return config.C.Database.Schema.App + ".tb_user"
}
