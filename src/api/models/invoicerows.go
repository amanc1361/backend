package models

import "gorm.io/gorm"

type InvoiceRow struct {
	gorm.Model
	InvoiceID  uint `json:"invoice_id"`
	ObjectId int `json:"object_id"`
	Name string `json:"name"`
	Unitname string  `json:"unit_name"`
	Count int   `json:"count"`
	Price int  `json:"price"`
	Discount int  `json:"discount"`

	




}