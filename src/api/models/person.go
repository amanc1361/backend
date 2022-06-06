package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name             string `json:"name" gorm:"uniqueIndex:idx_name;size:100;not null;"`
	Family           string ` json:"family" gorm:"uniqueIndex:idx_name;size:100;not null;"`
	TypePeople       int    `json:"type_people"`
	Tel              string `json:"tel"`
	Fax              string `json:"fax"`
	Mobile           string `json:"mobile"`
	Address          string `json:"address"`
	NationalNumber   string `json:"national_number"`
	EconomyNumber    string `json:"economy_number"`
	RegisterNumber   string `json:"register_number"`
	ShSh             string `json:"sh_sh"`
	InsurnanceNumber string `json:"insurnance_number"`
	National         string `json:"national"`
	PostalCode       string `json:"postal_code"`
	JobID            int    `json:"job_id"`
	BirthDate        string `json:"birth_date"`
	BrithRegister    string `json:"birth_register"`
	CompanyID        int    `gorm:"uniqueIndex:idx_name" json:"company_id"`
	DetailedName string `json:"detailed_name"`
	DetailedID       int    `json:"detailed_id"`
	IsMarried        int    `json:"is_married"`
	MarriedDate      string `json:"married_date"`
}
