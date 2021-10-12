package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryGroupCRUD struct {
	db *gorm.DB
}

func NewRepositoryGroupCRUD(db *gorm.DB) *repositoryGroupCRUD {

	return &repositoryGroupCRUD{db}
}

func (r *repositoryGroupCRUD) Save(group models.Group) (models.Group, error) {

	var err error

	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Create(&group).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return group, nil
	}

	return models.Group{}, err

}

func (r *repositoryGroupCRUD) FindAll(companyid int) ([]models.Group, error) {

	var err error
	// outresult:=[]myresult{}
	done := make(chan bool)
	groups := []models.Group{}

	go func(ch chan<- bool) {

		// result:=r.db.Model(&models.Group{}).Select("ledger.id,ledgers.name,ledgers.code,groups.id,groups.nature").Joins("left join groups on ledgers.groupid=groups.id").Scan(&outresult)
		result := r.db.Model(&models.Group{}).Where("company_id=?", companyid).Find(&groups)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return groups, nil
	}

	return nil, err

}

func (r *repositoryGroupCRUD) FindById(uid uint32) (models.Group, error) {

	var err error
	group := models.Group{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("id=?", uid).First(&group)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	

	if channels.Ok(done) {
		return group, nil
	}

	return models.Group{}, err

}

func (r *repositoryGroupCRUD) Update(group models.Group) (int64, error) {

	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {

		err = r.db.Save(&group).Error

		ch <- true

	}(done)

	if channels.Ok(done) {
		if err != nil {
			return 0, err

		}
	}

	return 1, nil

}

func (r *repositoryGroupCRUD) Delete(uid int32) (int64, error) {

	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {

		result = r.db.Where("id=?", uid).Delete(&models.Group{})

		ch <- true

	}(done)

	if channels.Ok(done) {
		if result.Error != nil {
			return 0, result.Error

		}
	}

	return result.RowsAffected, nil

}

func (r *repositoryGroupCRUD) FindByCode(uid uint, companyid int) (models.Group, error) {

	var err error
	group := models.Group{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		result := r.db.Where("Code=? and company_id=", uid).First(&group)
		err = result.Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return group, nil
	}

	return models.Group{}, err

}

func (r *repositoryGroupCRUD) GetLastCode(companyid int) (uint, error) {
	var code uint

	result := r.db.Raw("call getlastgroupcode(?)", companyid).Scan(&code)

	if result.Error != nil {

		return 0, result.Error
	}
	return code, nil
}
