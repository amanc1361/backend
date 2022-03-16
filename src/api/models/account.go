package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	AccountNumber   string `json:"account_number"`
	AccountType     int    `json:"account_type"`
	DetailedName    string `json:"detailed_name"`
	DetailedId      int    `json:"detailed_id"`
	Owner           string `json:"owner"`
	CreateSolarDate string `json:"create_solar_date"`
	IsActive        bool   `json:"is_active"`
	ShebaNumber     string `json:"sheba_number"`
	CardNumber      string `json:"card_number"`
	CompanyId 		int `json:"company_id"`
	YearId int `json:"year_id"`
}