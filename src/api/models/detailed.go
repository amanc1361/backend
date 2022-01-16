package models

import "gorm.io/gorm"

type Detailed struct {
	gorm.Model
	Name        string `json:"name" gorm:"uniqueIndex:idx_name;uniqueIndex:idx_code;"`
	Code        uint   `json:"code" gorm:"uniqueIndex:idx_code"`
	IsFloat     bool   `gorm:"default:true"`
	LedgerID    uint   `json:"ledger_id"`
	SubLedgerID uint   `json:"sub_ledger_id"`
	CompanyId   uint   `json:"company_id" gorm:"uniqueIndex:idx_name;uniqueIndex:idx_code;"`
}
