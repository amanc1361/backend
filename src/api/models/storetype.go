package models

import "gorm.io/gorm"

type StoreType struct {
	gorm.Model
	Name string `json:"name"`
	StoreId int  `json:"store_id"`
}