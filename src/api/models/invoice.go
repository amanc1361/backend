package models

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model
	PersonID         int `json:"peopleId"`
	InvoiceNumber    int `json:"invoiceNumber"`
	InvoiceBuynumber string `json:"buynumber"`
	SolarDate        string `json:"solarDate"`
	DueDate          string `json:"dueDate"`
	InvoiceTypeId      int `json:"invoiceTypeId"`
	Amount           int `json:"amount"`
	Description      string `json:"description"`
	Tax int `json:"tax"`
	ValueAdd int `json:"value_add"`
	ValueExtended int `json:"value_extended"`
	CompanyId int `json:"companyId"`
	YearId int  `json:"yearId"`
	InvoiceRows []InvoiceRow
}
