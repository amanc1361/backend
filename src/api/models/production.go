package models

import "gorm.io/gorm"

type Production struct {
	gorm.Model
	ProductId int `json:"product_id"`
	Amount float32 `json:"amount"`
	SolarFrom string `json:"solar_from"`
	SolarTo string `json:"solar_to"`
	
}