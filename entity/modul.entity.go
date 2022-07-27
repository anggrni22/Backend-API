package entity

import (
	"akrab-bangkit2022-api/config"
)

type Modul struct {
	ID     int    `json:"id"`
	Materi string `json:"materi"`
	Image  string `json:"image"`
	Description string `json:"description"`
}

func (Modul) TableName() string {
	return config.C.Database.Schema.App + ".tb_modul"
}
