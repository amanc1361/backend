package models

import "gorm.io/gorm"

type AccountType struct {
	gorm.Model
	Name string `json:"name"`
	CompanyId int `json:"company_id"`
	IsActive bool `json:"is_active"`
	
}