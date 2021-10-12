package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryCompanyType struct {
	db *gorm.DB
}

func NewRepositoryCompanyType(db *gorm.DB) *repositoryCompanyType {
	return &repositoryCompanyType{db}
}

func (r *repositoryCompanyType) FindAll() ([]models.CompanyType, error) {

	var err error
	companyTpes := []models.CompanyType{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Find(&companyTpes)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return companyTpes, nil
	}

	return nil, err

}

func (r *repositoryCompanyType) FindById(uid uint32) (models.CompanyType, error) {

	var err error
	companyType := models.CompanyType{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&companyType)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return companyType, nil
	}

	return models.CompanyType{}, err

}
