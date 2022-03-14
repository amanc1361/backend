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