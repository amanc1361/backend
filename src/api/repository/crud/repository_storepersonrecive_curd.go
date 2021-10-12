package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryStorePersonReciveCRUD struct {
	db *gorm.DB
}

func NewRepositoryStorePersonReciveCRUD(db *gorm.DB) *repositoryStorePersonReciveCRUD {

	return &repositoryStorePersonReciveCRUD{db}
}

func (r *repositoryStorePersonReciveCRUD) Save(storepersonrecive models.StorePersonRecive) (models.StorePersonRecive, error) {

	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Save(&storepersonrecive).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storepersonrecive, nil
	}

	return models.StorePersonRecive{}, err

}

func (r *repositoryStorePersonReciveCRUD) FindAll(companyid int, search string) ([]models.StorePersonRecive, error) {

	var err error
	storepersonreciveis := []models.StorePersonRecive{}
	done := make(chan bool)
	var result *gorm.DB

	go func(ch chan<- bool) {
		if len(search) == 0 {
			result = r.db.Model(&models.StorePersonRecive{}).Where("company_id=?", companyid).Find(&storepersonreciveis)
		} else {
			result = r.db.Model(&models.StorePersonRecive{}).Where("company_id=? and (name LIKE ? or  family like ?)", companyid, "%"+search+"%", "%"+search+"%").Find(&storepersonreciveis)
		}
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storepersonreciveis, nil
	}

	return nil, err

}

func (r *repositoryStorePersonReciveCRUD) FindById(uid uint32) (models.StorePersonRecive, error) {

	var err error
	storepersonrecive := models.StorePersonRecive{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&storepersonrecive)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storepersonrecive, nil
	}

	return models.StorePersonRecive{}, err

}

func (r *repositoryStorePersonReciveCRUD) Update(storepersonrecive models.StorePersonRecive) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Save(&storepersonrecive)

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryStorePersonReciveCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.StorePersonRecive{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
