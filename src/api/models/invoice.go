package models

import "gorm.io/gorm"

type Invocie struct {
	gorm.Model
	PersonID         int
	InvoiceNumber    int
	InvoiceBuynumber int
	SolarDate        string
	DueDate          string
	InvoiceType      int
	Amount           int
	Description      string
}
