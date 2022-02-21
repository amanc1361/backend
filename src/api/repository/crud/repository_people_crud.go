package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"
	"errors"

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
		// detailed, err = r1.Save(detailed)
		detailed.Code, err = r1.GetLastCode(int(detailed.CompanyId))
		err = tx.Create(&detailed).Error

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
		}
		ch <- true
	}(done)
if channels.Ok(done) {
		return People, tx.Commit().Error
	}
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
	var err error
	done := make(chan bool)
	detailed := models.Detailed{}
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if People.DetailedID != 0 {
		if People.TypePeople != 2 {
			detailed.Name = People.Name + " " + People.Family
		} else {
			detailed.Name = People.Name
		}
		detailed.ID = uint(People.DetailedID)
		detailed.CompanyId = uint(People.CompanyID)
		err := tx.Model(&models.Detailed{}).Where("id=?", detailed.ID).Update("name", detailed.Name).Error
		if err != nil {
			tx.Rollback()
			return 0, err

		}

	}

	go func(ch chan<- bool) {

		err = tx.Save(&People).Error
		if err != nil {
			tx.Rollback()
			ch <- false
		}
		ch <- true

	}(done)

	if channels.Ok(done) {

		return result.RowsAffected, tx.Commit().Error
	}

	return 0, err

}

func (r *repositoryPeopleCRUD) Delete(peopleid int32, detailedid int32) (int64, error) {
	var err error
	done := make(chan bool)
	var count int64
	go func(ch chan<- bool) {
		err = r.db.Model(&models.DocumentRows{}).Where("detailed_id=?", detailedid).Count(&count).Error
		if err != nil {
			ch <- false
		}
		if count == 0 {
			err = r.db.Where("id=?", peopleid).Delete(&models.Person{}).Error
			if err != nil {
				ch <- false
			}
			ch <- true
		} else {
			err = errors.New("این شخص یا شرکت گردش مالی دارد")
			ch <- false
		}

	}(done)

	if channels.Ok(done) {

		return 1, nil

	}

	return 0, err

}


func (r *repositoryPeopleCRUD) GetRemPerson(companyid int,yearid int,detailedid int, solardate string)(int,error) {
	var err error
	done:=make(chan bool)
	var rem int
	
	go func(ch chan<-bool) {
      err=r.db.Raw(`select sum(debtor)-sum(creditor) as rem from document_rows
	  join documents on documents.id=document_rows.document_id
	  where detailed_id=? and documents.deleted_at is null
	  and documents.company_id=? and documents.year_id=? and
	  documents.solar_date<=?`,detailedid,companyid,yearid,solardate).Take(&rem).Error
	  if err!=nil {
		  ch<-false
		  return
	  }
	  ch<-true
	}(done)

	if channels.Ok(done) {
		return 1,nil
	}
	return 0,err

}