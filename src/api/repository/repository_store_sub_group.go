package repository

import "back-account/src/api/models"

type StoreSubGroupRepository interface {
	Save(models.StoreSubGroup) (models.StoreSubGroup, error)
	FindAll() ([]models.StoreSubGroup, error)
	FindById(uint32) (models.StoreSubGroup, error)
	Update(models.StoreSubGroup) (int64, error)
	Delete(int32) (int64, error)
}
