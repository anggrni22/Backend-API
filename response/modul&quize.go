package response

import "akrab-bangkit2022-api/entity"

type AllRsp struct {
	Level 		string    		 `json:"level"`
	ImageLevel string `json:"image_level"`
	Tipe string `json:"tipe"`
	Modul []entity.Modul `json:"modul"`
	Quiz  []entity.Quiz  `json:"Quiz"`
}



