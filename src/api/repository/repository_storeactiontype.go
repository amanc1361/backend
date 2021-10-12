package repository

import "back-account/src/api/models"

type StoreActionType interface {
	Save(models.StoreActionType) (models.StoreActionType, error)
	FindAll(companyid int, search string) ([]models.StoreActionType, error)
	FindById(uint32) (models.StoreActionType, error)
	Update(models.StoreActionType) (int64, error)
	Delete(int32) (int64, error)
}
