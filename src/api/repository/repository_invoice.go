package repository

import "back-account/src/api/models"

type Inovice interface {
	Save( invoice models.Invocie)(models.Invocie,error)
	Delete(invoiceid int)(int,error)
	Update(invoice models.Invocie,invoiceid int)(models.Invocie,error)
	GetInvocie(invoiceid int)(models.Invocie,error)
	GetAll(companyid int,yeaid int)([]models.Invocie,error)
	GetLastInvoiceNumber(companyid int,yearid int,invoicetype int)(int,error)
	

}