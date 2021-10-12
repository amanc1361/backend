package models

import "gorm.io/gorm"

type DocumentType struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
}
