package entity

import "akrab-bangkit2022-api/config"

type Level struct {
	ID     	   int    `json:"id"`
	ImageLevel string `json:"image_level"`
	Level      string `json:"level"`
	Tipe 	   string `json:"tipe"`
}

func (Level) TableName() string {
	return config.C.Database.Schema.App + ".level"
}
