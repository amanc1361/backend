package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"
	"time"

	"gorm.io/gorm"
)

type repositoryYearCrud struct {
	db *gorm.DB
}

func NewRepositoryYearCrud(db *gorm.DB) *repositoryYearCrud {
	return &repositoryYearCrud{db}
}

func (r *repositoryYearCrud) Save(year models.Year) (models.Year, error) {

	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Create(&year).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return year, nil
	}

	return models.Year{}, err

}

func (r *repositoryYearCrud) FindAll(companyid int) ([]models.Year, error) {

	var err error
	Years := []models.Year{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Model(&models.Year{}).Where("company_id=?", companyid).Find(&Years)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return Years, nil
	}

	return nil, err

}

func (r *repositoryYearCrud) FindById(uid uint32) (models.Year, error) {

	var err error
	Year := models.Year{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&Year)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return Year, nil
	}

	return models.Year{}, err

}

func (r *repositoryYearCrud) Update(year models.Year) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	year.CreatedAt = time.Now()
	go func(ch chan<- bool) {

		result = r.db.Save(year)

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryYearCrud) Delete(uid uint32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.Year{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
