package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryUsersCRUD struct {
	db *gorm.DB
}

func NewRepositoryUsersCRUD(db *gorm.DB) *repositoryUsersCRUD {

	return &repositoryUsersCRUD{db}
}

func (r *repositoryUsersCRUD) Save(user models.User) (models.User, error) {

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

	return models.User{}, err

}

func (r *repositoryUsersCRUD) FindAll() ([]models.User, error) {

	var err error
	users := []models.User{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Find(&users)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return users, nil
	}

	return nil, err

}

func (r *repositoryUsersCRUD) FindById(uid uint32) (models.User, error) {

	var err error
	user := models.User{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&user)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return user, nil
	}

	return models.User{}, err

}

func (r *repositoryUsersCRUD) Update(uid int32, user models.User) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Model(&models.User{}).Where("id=?", uid).Updates(models.User{UserName: user.UserName, Password: user.Password, RoleId: user.RoleId})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryUsersCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.User{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
