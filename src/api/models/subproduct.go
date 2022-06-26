package models

import "gorm.io/gorm"

type SubProduct struct {
	gorm.Model
	Name string `json:"name"`
	ProductID int `json:"procuct_id"`
	StoreObjectId int `json:"store_object_id"`
	
}