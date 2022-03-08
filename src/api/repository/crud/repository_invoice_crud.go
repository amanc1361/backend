package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
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
invoicenumber,err:=r.GetInvoiceNumber(invoice.CompanyId,invoice.YearId,invoice.InvoiceTypeId)
if err!=nil {
	return models.Invoice{},err
}

invoice.InvoiceNumber=invoicenumber
go func(ch chan<-bool) {
    err:=r.db.Create(&invoice).Error
	if err!=nil {
		ch<-false 
		return 
	}
	ch<-true
}(done)

if channels.Ok(done) {
	r.db.Model(&models.StoreAction{}).Where("id=?",invoice.StoreId).Update("invoice_id",invoice.ID)
	return invoice,err
}
return models.Invoice{},err

}



func (r *invoiceRepository) GetAll(companyid int,yearid int,invoicetype int)([]modelsout.Invoice,error) {
	var err error
	 invoices :=[]modelsout.Invoice{}
	 done:=make(chan bool)
	go func(ch chan<-bool) {
		err=r.db.Table("invoices").
		Select("invoices.id,invoices.solar_date,invoices.invoice_number,invoices.amount,invoices.description,concat( people.name,' ',people.family) as customer_name ").
		Joins("join people on invoices.person_id=people.id").
		Where("invoices.company_id=? and invoices.year_id=? and invoices.invoice_type_id=?",companyid,yearid,invoicetype).
		Order("invoices.invoice_number").
		Scan(&invoices).Error
		
		if err!=nil {
			ch<-false
			return
		}
		ch<-true
	}(done)
	if channels.Ok(done) {
		return invoices,err
	}
	return []modelsout.Invoice{},err
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
	invoicenumber:=0
	done:=make(chan bool)
	go func(ch chan<-bool) {
		err=r.db.Raw("select max(invoice_number) from invoices where company_id=? and year_id=? and invoice_type_id=?",
	                     companyid,yearid,invoicetypeid ).Take(&invoicenumber).Error
		  
	    if err!=nil {
			ch<-false
			return
		}
		ch<-true
	    
		}(done)
	if channels.Ok(done) {
		 return invoicenumber+1,nil
	}	
	 if err==gorm.ErrRecordNotFound {
		  return 1,nil
	 }
	return 1,nil
}
