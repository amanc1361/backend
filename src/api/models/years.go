package models

import "gorm.io/gorm"

type Year struct {
	gorm.Model
	Name      string `gorm:"not null;unique"`
	CompanyId uint
	IsOpen    bool `gorm:"default:false"`
}
