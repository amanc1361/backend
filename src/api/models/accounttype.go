package models

import "gorm.io/gorm"

type AccountType struct {
	gorm.Model
	Name string `json:"name"`
	IsActive bool `json:"is_active"`
	
}