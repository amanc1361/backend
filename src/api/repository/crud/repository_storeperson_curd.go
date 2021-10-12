package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryStorePersonCRUD struct {
	db *gorm.DB
}

func NewRepositoryStorePersonCRUD(db *gorm.DB) *repositoryStorePersonCRUD {

	return &repositoryStorePersonCRUD{db}
}

func (r *repositoryStorePersonCRUD) Save(storeperson models.StorePerson) (models.StorePerson, error) {

	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Save(&storeperson).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeperson, nil
	}

	return models.StorePerson{}, err

}

func (r *repositoryStorePersonCRUD) FindAll(companyid int, search string) ([]models.StorePerson, error) {

	var err error
	storepersonis := []models.StorePerson{}
	done := make(chan bool)
	var result *gorm.DB

	go func(ch chan<- bool) {
		if len(search) == 0 {
			result = r.db.Model(&models.StorePerson{}).Where("company_id=?", companyid).Find(&storepersonis)
		} else {
			result = r.db.Model(&models.StorePerson{}).Where("company_id=? and (name LIKE ? or  family like ?)", companyid, "%"+search+"%", "%"+search+"%").Find(&storepersonis)
		}
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storepersonis, nil
	}

	return nil, err

}

func (r *repositoryStorePersonCRUD) FindById(uid uint32) (models.StorePerson, error) {

	var err error
	storeperson := models.StorePerson{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&storeperson)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return storeperson, nil
	}

	return models.StorePerson{}, err

}

func (r *repositoryStorePersonCRUD) Update(storeperson models.StorePerson) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Save(&storeperson)

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryStorePersonCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.StorePerson{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
