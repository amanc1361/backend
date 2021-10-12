package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryStoreActionTypeCRUD struct {
	db *gorm.DB
}

func NewRepositoryStoreActionTypeCRUD(db *gorm.DB) *repositoryStoreActionTypeCRUD {

	return &repositoryStoreActionTypeCRUD{db}
}

func (r *repositoryStoreActionTypeCRUD) Save(storeactiontype models.StoreActionType) (models.StoreActionType, error) {

	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Save(&storeactiontype).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeactiontype, nil
	}

	return models.StoreActionType{}, err

}

func (r *repositoryStoreActionTypeCRUD) FindAll(companyid int, search string) ([]models.StoreActionType, error) {

	var err error
	storeactiontypeis := []models.StoreActionType{}
	done := make(chan bool)
	var result *gorm.DB

	go func(ch chan<- bool) {
		if len(search) == 0 {
			result = r.db.Model(&models.StoreActionType{}).Where("company_id=?", companyid).Find(&storeactiontypeis)
		} else {
			result = r.db.Model(&models.StoreActionType{}).Where("company_id=? and name LIKE  ?", companyid, "%"+search+"%").Find(&storeactiontypeis)
		}
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeactiontypeis, nil
	}

	return nil, err

}

func (r *repositoryStoreActionTypeCRUD) FindById(uid uint32) (models.StoreActionType, error) {

	var err error
	storeactiontype := models.StoreActionType{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&storeactiontype)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeactiontype, nil
	}

	return models.StoreActionType{}, err

}

func (r *repositoryStoreActionTypeCRUD) Update(storeactiontype models.StoreActionType) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Save(&storeactiontype)

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryStoreActionTypeCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.StoreActionType{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
