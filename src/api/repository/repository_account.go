package repository

import "back-account/src/api/models"
type AccountRepository interface {
	Save(account models.Account)(models.Account,error)
	// Update(account models.Account)(models.Account,error)
	// Delete(accountid int)error
	// Gets(companyid int)([]models.Account,error)
	// Get(accountid int)(models.Account,error)
	}