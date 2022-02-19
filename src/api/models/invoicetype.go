package models

import (
	"gorm.io/gorm"
)

type InvoiceType struct {
	gorm.Model
	Name string `json:"name" gorm:"primaryKey;not null"`
	LedgerId int 	`json:"ledger_id"`
	SubLedgerId int  `json:"sub_ledger_id"`
	CompanyID int  `json:"company_id"`
	

}