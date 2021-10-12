package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"
	"fmt"

	"gorm.io/gorm"
)

type repositoryPeopleCRUD struct {
	db *gorm.DB
}

func NewRepositoryPeopleCRUD(db *gorm.DB) *repositoryPeopleCRUD {

	return &repositoryPeopleCRUD{db}
}

func (r *repositoryPeopleCRUD) Save(People models.Person) (models.Person, error) {

	var err error

	done := make(chan bool)

	detailed := models.Detailed{}
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("error in herer")
			tx.Rollback()
		}
	}()

	if People.DetailedID == 0 {
		if People.TypePeople != 2 {
			detailed.Name = People.Name + " " + People.Family
		} else {
			detailed.Name = People.Name
		}
		detailed.CompanyId = uint(People.CompanyID)
		r1 := NewRepositoryDetailedCRUD(r.db)
		detailed, err = r1.Save(detailed)

		if err != nil {
			tx.Rollback()
			return models.Person{}, err

		}
		People.DetailedID = int(detailed.ID)

	}

	go func(ch chan<- bool) {

		err = tx.Create(&People).Error

		if err != nil {

			ch <- false
			tx.Rollback()
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return People, tx.Commit().Error
	}
	fmt.Printf("5")
	return models.Person{}, err

}

func (r *repositoryPeopleCRUD) FindAll(companyid int, search string) ([]models.Person, error) {

	var err error
	Peopleis := []models.Person{}
	done := make(chan bool)
	var result *gorm.DB

	go func(ch chan<- bool) {
		if len(search) == 0 {
			result = r.db.Model(&models.Person{}).Where("company_id=?", companyid).Find(&Peopleis)
		} else {
			result = r.db.Model(&models.Person{}).Where("company_id=? and (name LIKE ? or  family like ?)", companyid, "%"+search+"%", "%"+search+"%").Find(&Peopleis)
		}
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return Peopleis, nil
	}

	return nil, err

}

func (r *repositoryPeopleCRUD) FindById(uid uint32) (models.Person, error) {

	var err error
	People := models.Person{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&People)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return People, nil
	}

	return models.Person{}, err

}

func (r *repositoryPeopleCRUD) Update(People models.Person) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Save(&People)

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryPeopleCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.Person{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}
