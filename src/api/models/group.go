package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name      string `gorm:"primaryKey;not null"`
	Nature    byte
	Code      uint `gorm:"primaryKey;not null"`
	CompanyId uint `gorm:"primaryKey;not null"`
	Ledger    []Ledger
}
