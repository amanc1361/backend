package models

import "gorm.io/gorm"

type CostCenter struct {
	gorm.Model
	CompanyId int `json:"company_id"`
	LedgerId int `json:"ledger_id"`
	SubLedgerId int `json:"sub_ledger_id"`
	DetailedId int `json:"detailed_id"`
	PriceName string `json:"price_name"`

	
}