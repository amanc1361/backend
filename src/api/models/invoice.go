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
	
	CompanyId int `json:"companyId"`
	YearId int  `json:"yearId"`
	InvoiceRows []InvoiceRow
}
