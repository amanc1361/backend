package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	DocumentNumber uint `gorm:"not null"`
	Amount         uint
	YearID         uint `gorm:"not null"`
	UserID         uint `gorm:"not null"`
	CompanyID      uint `gorm:"not null"`
	SolarDate      string
	DocumentTypeID uint `gorm:"not null"`
	Description    string
	Istemp         bool `gorm:"default:true"`
	DocumentRows   []DocumentRows
}
