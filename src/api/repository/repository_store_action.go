package repository

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
)

type StoreActionRepository interface {
	Save(models.StoreAction) (models.StoreAction, error)
	FindAll(search string, tpeaction int, yearid int, companyid int, storeid int) ([]models.StoreAction, error)
	FindById(uint32) (models.StoreAction, error)
	Update(models.StoreAction) (int64, error)
	Delete(int32) (int64, error)
	GetCountObject(companyid int, yearid int, id int) (float32, error)
	Getsend(companyid int, yearid int) ([]modelsout.StoreActiondoc, error)
	Getrecive(companyid int, yearid int) ([]modelsout.StoreActiondoc, error)
	GetStoreActionRows(uint32) ([]models.StoreActionRow, error)
	GetKardex(companyid int, yearid int, id int) ([]modelsout.Kradex, error)
	GetRemObject(companyid int, yearid int, id int) ([]modelsout.Remobject, error)
	GetPriceObject(companyid int, yearid int, objectid int, solardate string) (float64, error)
        GetStoreIdbyDocumentID(int)(int,int,error)
}
