package repository

import "back-account/src/api/models"

type StoreTypeRepository interface {
	Save(models.StoreType) (models.StoreType,error)
	Update(models.StoreType)(int64,error)
	Delete(int32)(int32,error)
	FindAll(storeid int)([]models.StoreType,error)
	GetById(storetypeid int)(models.StoreType,error)
	

}