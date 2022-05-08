package models

import "gorm.io/gorm"

type WareHouseingRows struct {
	gorm.Model
	Code int    `json:"code"`
	Name string `json:"name"`
	Rem             int    `json:"rem"`
	Conflict        int    `json:"conflict"`
}