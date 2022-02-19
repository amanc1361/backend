package repository

import "back-account/src/api/models"

type Inovice interface {
	Save( invoice models.Invoice)(models.Invoice,error)
	// Delete(invoiceid int)(int,error)
	// Update(invoice models.Invocie,invoiceid int)(models.Invocie,error)
	GetInvocie(invoiceid int)(models.Invoice,error)
	GetAll(companyid int,yeaid int,invoicetype int)([]models.Invoice,error)
	GetLastInvoiceNumber(companyid int,yearid int,invoicetype int)(int,error)
	GetInovicTypies(companyid int) ([]models.InvoiceType,error)
	

}