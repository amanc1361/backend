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

func (r *invoiceRepository) Save(invoice models.Invoice) (models.Invoice,error) {
var err error
done:=make(chan bool)
invoicenumber,err:=r.GetLastInvoiceNumber(invoice.CompanyId,invoice.YearId,2)
if err!=nil {
	return models.Invoice{},err
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
return models.Invoice{},err

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

func (r *invoiceRepository) GetAll(companyid int,yearid int,invoicetype int)([]models.Invoice,error) {
	var err error
	 invoices :=[]models.Invoice{}
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
	return []models.Invoice{},err
}

func (r *invoiceRepository) GetInvocie(invoiceid int)(models.Invoice,error) {
	var err error
	var invoice models.Invoice
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
	return models.Invoice{},err
}
func (r *invoiceRepository) GetSellTypeis(companyid int) ([]models.SellType,error) {
	var err error 
	invoicetypies:=[]models.SellType{}
	done:=make(chan bool)

	go func(ch chan<-bool) {
		err=r.db.Where("company_id=? ",companyid).Find(&invoicetypies).Error
		if err!=nil {
			ch<-false 
			return
		}
		ch<-true
	}(done)
	if channels.Ok(done) {
		return invoicetypies,err
	}
	return nil,err
}

func (r *invoiceRepository) GetInvoiceNumber(companyid int,yearid int,invoicetypeid int)(int ,error) {
	var err error 
	var invoicenumber int
	invoicenumber=1
	done:=make(chan bool)
	go func(ch chan<-bool) {
		  err= r.db.Model(models.Invoice{}).Where("company_id=? and year_id=? and invoice_type_id=?",companyid,yearid,invoicetypeid).Select("invoice_number").Take(&invoicenumber).Error
	    if err!=nil {
			ch<-false
			return
		}
		ch<-true
	    
		}(done)
	if channels.Ok(done) {
		 return invoicenumber,err
	}	
	 if err==gorm.ErrRecordNotFound {
		  return 1,nil
	 }
	return 0,err
}
