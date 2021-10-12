package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryUnitCRUD struct {
	db *gorm.DB
}

func NewRepositoryUnitCRUD(db *gorm.DB) *repositoryUnitCRUD {

	return &repositoryUnitCRUD{db}
}

func (r *repositoryUnitCRUD) Save(unit models.Unit) (models.Unit, error) {

	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Create(&unit).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return unit, nil
	}

	return models.Unit{}, err

}

func (r *repositoryUnitCRUD) FindAll(companyid int, search string) ([]models.Unit, error) {

	var err error
	unitis := []models.Unit{}
	done := make(chan bool)
	var result *gorm.DB
	go func(ch chan<- bool) {
		if len(search) == 0 {
			result = r.db.Model(&models.Unit{}).Where("company_id=?", companyid).Find(&unitis)
		} else {
			result = r.db.Model(&models.Unit{}).Where("company_id=? and name LIKE ? )", companyid, "%"+search+"%").Find(&unitis)
		}
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return unitis, nil
	}

	return nil, err

}

func (r *repositoryUnitCRUD) FindById(uid uint32) (models.Unit, error) {

	var err error
	unit := models.Unit{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&unit)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return unit, nil
	}

	return models.Unit{}, err

}

func (r *repositoryUnitCRUD) Update(unit models.Unit) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Save(&unit)

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryUnitCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.Unit{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
