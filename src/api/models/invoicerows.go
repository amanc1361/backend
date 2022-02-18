package models

import "gorm.io/gorm"

type InvoiceRow struct {
	gorm.Model
	InvoiceId int `json:"invoiceId"`
	ObjectID int `json:"objectId"`
	Count int  `json:"count"`
	Price int 	`json:"price"`
	Discount int `json:"discount"`
	Tax int `json:"tax"`
	ValueAdd int `json:"vlaueAdd"`
	




}