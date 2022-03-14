package repository

import "back-account/src/api/models"
type AccountRepository interface {
	Save(account models.Account)(models.Account,error)
	Update(account models.Account)(models.Account,error)
	Delete(accountid int)error
	GetAll(companyid int)([]models.Account,error)
	GetByID(accountid int)(models.Account,error)
	}