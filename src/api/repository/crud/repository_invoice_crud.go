package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvocieRepository(db *gorm.DB) *invoiceRepository {
	return &invoiceRepository{db}
}

func (r *invoiceRepository) Save(invoice models.Invocie) (models.Invocie,error) {
var err error
done:=make(chan bool)
invoicenumber,err:=r.GetLastInvoiceNumber(invoice.CompanyId,invoice.YearId,2)
if err!=nil {
	return models.Invocie{},err
}

invoice.InvoiceNumber=invoicenumber+1
go func(ch chan<-bool) {
    err:=r.db.Create(&invoice).Error
	if err!=nil {
		ch<-false 
		return 
	}
	ch<-true
}(done)

if channels.Ok(done) {
	return invoice,err
}
return models.Invocie{},err

}

func (r *invoiceRepository ) GetLastInvoiceNumber(companyid int,yearid int,invoicetype int)(int,error) {
	var err error 
	done:=make(chan bool)
	var invoicenumber int

	go func(ch chan<-bool) {
		err=r.db.Select("invoiceNumber").Where("companyId=? and year_id=? and invoiceType=?",companyid,yearid,invoicetype).Take(&invoicenumber).Error
		if err!=nil {
			ch<-false
			return
		}
		ch<-true
	}(done)
	if channels.Ok(done) {
		return invoicenumber,err
	}
	return 0,err
}

func (r *invoiceRepository) GetAll(companyid int,yearid int,invoicetype int)([]models.Invocie,error) {
	var err error
	 invoices :=[]models.Invocie{}
	 done:=make(chan bool)
	go func(ch chan<-bool) {
		err=r.db.Where("companyId=? and yearId=? and invoiceType=?",companyid,yearid,invoicetype).Find(&invoices).Error
		if err!=nil {
			ch<-false
			return
		}
		ch<-true
	}(done)
	if channels.Ok(done) {
		return invoices,err
	}
	return []models.Invocie{},err
}

func (r *invoiceRepository) GetInvocie(invoiceid int)(models.Invocie,error) {
	var err error
	var invoice models.Invocie
	done:=make(chan bool)
	go func(ch chan<-bool) {
		err=r.db.Where("id=?",invoiceid).Take(&invoice).Error
		if err!=nil {
			ch<-false
			return
		}
		ch<-true
	}(done)
	if channels.Ok(done) {
		return invoice,err
	}
	return models.Invocie{},err
}
