package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name          string `gorm:"unique;not null"`
	Image         string
	CompanyTypeID uint
	RegisterCode string `json:"register_code"`
	EconomyCode string `json:"economy_code"`
	NationalCode string `json:"national_code"`
	Address string `json:"address"`
	PsotalCode string `json:"postal_code"`
	Tel string `json:"tel"`
	Mobile string `json:"mobile"`
	Fax string `json:"fax"`
	Accounts []Account `json:"accounts"`

}
