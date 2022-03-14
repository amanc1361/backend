package models

import "gorm.io/gorm"

type BankBranch struct {
	gorm.Model
	Name string `json:"name"`
	BankID int `json:"bank_id"`
	BranchCode int `json:"branch_code"`
	CityId int `json:"city_id"`
}
