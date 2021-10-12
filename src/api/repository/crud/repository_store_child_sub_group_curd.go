package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryStoreChildSubGroupCRUD struct {
	db *gorm.DB
}

func NewRepositoryStoreChildSubGroupCRUD(db *gorm.DB) *repositoryStoreChildSubGroupCRUD {

	return &repositoryStoreChildSubGroupCRUD{db}
}

func (r *repositoryStoreChildSubGroupCRUD) Save(storechildsubgroup models.StoreChildSubGroup) (models.StoreChildSubGroup, error) {

	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Create(&storechildsubgroup).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storechildsubgroup, nil
	}

	return models.StoreChildSubGroup{}, err

}

func (r *repositoryStoreChildSubGroupCRUD) FindAll() ([]models.StoreChildSubGroup, error) {

	var err error
	storechildsubgroupis := []models.StoreChildSubGroup{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Find(&storechildsubgroupis)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storechildsubgroupis, nil
	}

	return nil, err

}

func (r *repositoryStoreChildSubGroupCRUD) FindById(uid uint32) (models.StoreChildSubGroup, error) {

	var err error
	storechildsubgroup := models.StoreChildSubGroup{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&storechildsubgroup)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storechildsubgroup, nil
	}

	return models.StoreChildSubGroup{}, err

}

func (r *repositoryStoreChildSubGroupCRUD) Update(storechildsubgroup models.StoreChildSubGroup) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Save(&storechildsubgroup)

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryStoreChildSubGroupCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.StoreChildSubGroup{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
