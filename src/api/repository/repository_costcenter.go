package repository

import "back-account/src/api/models"

type CostCenterRepository interface {
	Save(models.CostCenter)  (models.CostCenter,error)
	FindAll(companyid int) ([]models.CostCenter,error)
	FindByID(costcenterid int)(models.CostCenter,error)
	Delete(costcenterid int)(int,error)
	Update(models.CostCenter)(models.CostCenter,error)
	

}