package repository

import "back-account/src/api/models"

type WareHousingRepository interface {
	Save(models.WareHousing) (models.WareHousing, error)
	Update(models.WareHousing) (int64, error)
	FindAll(companyid int, yearid int) ([]models.WareHousing, error)
	Delete(int32) (int64, error)
	GetWareHousingById(int32) (models.WareHousing, error)
}