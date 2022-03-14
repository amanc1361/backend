package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryAccountCrud   struct {
	db *gorm.DB
}

func NewRepositoryAccountCRUD(db *gorm.DB) *repositoryAccountCrud {
	 return &repositoryAccountCrud{db}
} 

func (r *repositoryAccountCrud) Save(account models.Account)(models.Account,error ) {
	var err error
    done:=make(chan bool)

	go func(ch chan<-bool) {
		 err=r.db.Create(&account).Error
		 if err!=nil {
			 ch<-false
			 return
		 }
		 ch<-true
	}(done)
	if channels.Ok(done) {
		return account,nil

	}
	return models.Account{},err
}

func (r *repositoryAccountCrud) Update(account models.Account) (models.Account ,error) {
	var err error
    done:=make(chan bool)

	go func(ch chan<-bool) {
		 err=r.db.Updates(&account).Error
		 if err!=nil {
			 ch<-false
			 return
		 }
		 ch<-true
	}(done)
	if channels.Ok(done) {
		return account,nil

	}
	return models.Account{},err
	  
}

func (r *repositoryAccountCrud) Delete(accountid int)error {
	var err error
    done:=make(chan bool)

	go func(ch chan<-bool) {
		 err=r.db.Delete(models.Account{},accountid).Error
		 if err!=nil {
			 ch<-false
			 return
		 }
		 ch<-true
	}(done)
	if channels.Ok(done) {
		return nil

	}
	return err
}
func (r *repositoryAccountCrud) Get(accountid int) (models.Account,error) {
	var err error
    done:=make(chan bool)
	account:=models.Account{}

	go func(ch chan<-bool) {
		 err=r.db.Model(models.Account{}).Where("id=?",accountid).First(&account).Error
		 if err!=nil {
			 ch<-false
			 return
		 }
		 ch<-true
	}(done)
	if channels.Ok(done) {
		return account,nil

	}
	return models.Account{},err

}
func (r *repositoryAccountCrud) Gets()([]models.Account,error) {
	var err error
    done:=make(chan bool)
	acounts:=[]models.Account{}

	go func(ch chan<-bool) {
		 err=r.db.Take(&acounts).Error
		 if err!=nil {
			 ch<-false
			 return
		 }
		 ch<-true
	}(done)
	if channels.Ok(done) {
		return acounts,nil

	}
	return []models.Account{},err
}
