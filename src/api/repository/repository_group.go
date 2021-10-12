package repository

import "back-account/src/api/models"

type GroupRepository interface {
	Save(models.Group) (models.Group, error)
	FindAll(companyid int) ([]models.Group, error)
	FindById(uint32) (models.Group, error)
	FindByCode(id uint, companyid int) (models.Group, error)
	Update(models.Group) (int64, error)
	Delete(int32) (int64, error)
	GetLastCode(companyid int) (uint, error)
}
