package models

import (
	"gorm.io/gorm"
)

type InvoiceType struct {
	gorm.Model
	Name string `json:"name" gorm:"primaryKey;not null"`
	LedgerId int 	`json:"ledgerId"`
	SubLedgerId int  `json:"subLedgerId"`

}