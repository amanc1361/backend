package models

import "gorm.io/gorm"

type City struct {
	gorm.Model
	Name string `json:"name"`
	CompanyId int `json:"company_id"`
}