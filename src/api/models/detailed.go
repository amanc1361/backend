package models

import "gorm.io/gorm"

type Detailed struct {
	gorm.Model
	Name        string
	Code        uint
	IsFloat     bool `gorm:"default:true"`
	LedgerID    uint
	SubLedgerID uint
	CompanyId   uint
}
