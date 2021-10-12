package models

import "gorm.io/gorm"

type Ledger struct {
	gorm.Model
	Name      string `gorm:"uniqueIndex:idx_name"`
	Code      uint   `gorm:"uniqueIndex:idx_name"`
	Nature    byte
	GroupID   uint
	CompanyId uint `gorm:"uniqueIndex:idx_name"`
	SubLedger SubLedger
}
