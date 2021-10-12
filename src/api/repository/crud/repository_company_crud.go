package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryCompanyCRUD struct {
	db *gorm.DB
}

func NewRepositoryCompanyCRUD(db *gorm.DB) *repositoryCompanyCRUD {

	return &repositoryCompanyCRUD{db}
}

func (r *repositoryCompanyCRUD) Save(user models.Company) (models.Company, error) {

	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return user, nil
	}

	return models.Company{}, err

}

func (r *repositoryCompanyCRUD) FindAll(user_id int) ([]models.Company, error) {

	var err error
	companies := []models.Company{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Raw("SELECT * FROM companies where id in (select company_id from uesr_related_companies where user_id=?) ", user_id).Scan(&companies)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return companies, nil
	}

	return nil, err

}

func (r *repositoryCompanyCRUD) FindById(uid uint32) (models.Company, error) {

	var err error
	company := models.Company{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&company)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return company, nil
	}

	return models.Company{}, err

}

func (r *repositoryCompanyCRUD) Update(uid int32, company models.Company) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Model(&models.Company{}).Where("id=?", uid).Updates(models.Company{Name: company.Name, Image: company.Image, CompanyTypeID: company.CompanyTypeID})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryCompanyCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.Company{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
