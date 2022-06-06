package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryStoreTypeCRUD struct {
	db *gorm.DB
}

func NewRepositoryStoreTypeCrud(db *gorm.DB) *repositoryStoreTypeCRUD {
	return &repositoryStoreTypeCRUD{db}
}

func (r *repositoryStoreTypeCRUD) Save(storeType models.StoreType)(models.StoreType,error) {
	var err error
	done :=make(chan bool)
	go func(ch chan<-bool) {
		err=r.db.Save(&storeType).Error
		if err!=nil {
			ch<-false
			return
		}
		ch<-true
	}(done)
	if channels.Ok(done) {
		 return storeType,nil
	}
	return models.StoreType{},err
}