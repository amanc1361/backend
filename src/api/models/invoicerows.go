package models

import "gorm.io/gorm"

type InvoiceRow struct {
	gorm.Model
	InvoiceID  uint
	ObjectId int 
	Count int  
	Price int 
	Discount int 
	Tax int 
	ValueAdd int 
	




}