package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name             string `gorm:"uniqueIndex:idx_name;size:100;not null;"`
	Family           string `gorm:"uniqueIndex:idx_name;size:100;not null;"`
	TypePeople       int    `json:"type_people"`
	Tel              string
	Fax              string
	Mobile           string
	Address          string
	NationalNumber   string `json:"national_number"`
	EconomyNumber    string `json:"economy_number"`
	RegisterNumber   string `json:"register_number"`
	ShSh             string `json:"sh_sh"`
	InsurnanceNumber string `json:"insurnance_number"`
	National         string
	PostalCode       string `json:"postal_code"`
	JobID            int    `json:"job_id"`
	BirthDate        string `json:"birth_date"`
	BrithRegister    string `json:"birth_register"`
	CompanyID        int    `gorm:"uniqueIndex:idx_name" json:"company_id"`
	DetailedID       int    `json:"detailed_id"`
	IsMarried        int    `json:"is_married"`
	MarriedDate      string `json:"married_date"`
}
