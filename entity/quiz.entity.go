package entity

import "akrab-bangkit2022-api/config"

type Quiz struct {
	ID     int    `json:"id"`
	Soal   string `json:"soal"`
	Jawaban string `json:"jawaban"`
}

func (Quiz) TableName() string {
	return config.C.Database.Schema.App + ".tb_quiz"
}
