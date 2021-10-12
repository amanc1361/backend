package models

import "gorm.io/gorm"

type SubLedger struct {
	gorm.Model
	Name          string
	OwnerDetailed bool `gorm:"default:false"`
	Code          uint
	LedgerID      uint
	CompanyId     uint
	Nature        byte
}
