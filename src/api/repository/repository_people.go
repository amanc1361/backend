package repository

import "back-account/src/api/models"

type PeopleRepository interface {
	Save(models.Person) (models.Person, error)
	FindAll(companyid int, search string) ([]models.Person, error)
	FindById(uint32) (models.Person, error)
	Update(models.Person) (int64, error)
	GetLikeName(name string,companyid int)([]models.Detailed,error)
	Delete(pepopleid int32, detailedid int32) (int64, error)
	GetRemPerson(companyid int,yearid int,detailedid int, solardate string)(int,error)
}
