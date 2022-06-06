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

func (r *repositoryStoreTypeCRUD) Update(storeType models.StoreType)(int,error) {
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
		 return 1,nil
	}
	return 0,err
}
func (r *repositoryStoreTypeCRUD) FindAll(storeid int32)([]models.StoreType,error) { 
	var err error
	done :=make(chan bool)
	storetypies:=[]models.StoreType{}
	go func(ch chan<-bool) {
		err=r.db.Model(models.StoreType{}).Where("store_id=?",storeid).Find(storetypies).Error
		if err!=nil {
			ch<-false
			return
		}
		ch<-true
	}(done)
	if channels.Ok(done) {
		 return storetypies,nil
	}
	return []models.StoreType{},err
}
func (r *repositoryStoreTypeCRUD) GetById(storetypeid int32)(models.StoreType,error) { 
	var err error
	done :=make(chan bool)
	storetypie:=models.StoreType{}
	go func(ch chan<-bool) {
		err=r.db.Model(models.StoreType{}).Where("id",storetypeid).First(storetypie).Error
		if err!=nil {
			ch<-false
			return
		}
		ch<-true
	}(done)
	if channels.Ok(done) {
		 return storetypie,nil
	}
	return models.StoreType{},err
}
func (r *repositoryStoreTypeCRUD) Delete(storetypeid int32)(int32,error) { 
	var err error
	done :=make(chan bool)
	storetypie:=models.StoreType{}
	go func(ch chan<-bool) {
		err=r.db.Model(models.StoreType{}).Where("id",storetypeid).Delete(storetypie).Error
		if err!=nil {
			ch<-false
			return
		}
		ch<-true
	}(done)
	if channels.Ok(done) {
		 return 1,nil
	}
	return 0,err
}

