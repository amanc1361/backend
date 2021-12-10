package models

import "gorm.io/gorm"

type DocumentMerge struct {
	gorm.Model
	DocumentNumber uint `gorm:"not null"`
	Amount         uint
	YearID         uint `gorm:"not null"`
	UserID         uint `gorm:"not null"`
	CompanyID      uint `gorm:"not null"`
	SolarDate      string
	DocumentTypeID uint `gorm:"not null"`
	Description    string
	SolarFrom      string `json:"solar_from"`
	SolarTo        string `json:"solar_to"`
	DocFrom        int    `json:"doc_from"`
	DocTo          int    `json:"doc_to"`
}
