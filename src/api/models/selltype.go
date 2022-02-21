package models

import "gorm.io/gorm"

type SellType struct {
	gorm.Model
	Name string `json:"name"`
	LedgerId int 	`json:"ledger_id"`
	SubLedgerId int  `json:"sub_ledger_id"`
	CompanyID int  `json:"company_id"`
	Percent float32 `json:"percent"`


}
