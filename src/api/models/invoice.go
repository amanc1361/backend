package models

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model
	PersonID         int `json:"people_id"`
	InvoiceNumber    int `json:"invoice_number"`
	InvoiceBuynumber string `json:"buy_number"`
	SolarDate        string `json:"solar_date"`
	DueDate          string `json:"due_date"`
	InvoiceTypeId      int `json:"invoice_type_id"`
	SellTypeId int `json:"sell_type_id"`
	Amount           int `json:"amount"`
	Description      string `json:"description"`
	Tax int `json:"tax"`
	ValueAdd int `json:"value_add"`
	ValueExtended int `json:"value_extended"`
	CompanyId int `json:"company_id"`
	YearId int  `json:"year_id"`
	InvoiceRows []InvoiceRow `json:"invoice_rows"`
}
