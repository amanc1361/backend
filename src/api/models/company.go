package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name          string `gorm:"unique;not null"`
	Image         string
	CompanyTypeID uint
}
