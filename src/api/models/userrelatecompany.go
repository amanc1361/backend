package models

import "gorm.io/gorm"

type UesrRelatedCompany struct {
	gorm.Model
	UserID    int `json:"user_id"`
	CompanyID int `json:"company_id"`
}
