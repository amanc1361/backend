package repository

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
)

type Inovice interface {
	Save( invoice models.Invoice)(models.Invoice,error)
	// Delete(invoiceid int)(int,error)
	// Update(invoice models.Invocie,invoiceid int)(models.Invocie,error)
	GetInvocie(invoiceid int)(models.Invoice,error)
	GetAll(companyid int,yeaid int,invoicetype int)([]modelsout.Invoice,error)
	GetSellTypeis(companyid int) ([]models.SellType,error)
	GetInvoiceNumber(companyid int,yearid int,typeinvoice int)(int,error)
	

}