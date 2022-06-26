package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `json:"name"`
	UnitId int `json:"unit_id"`
	
}