package crud

type repositoryWarehousingCRUD struct {
	db *gorm.DB
}

func NewRepositoryWareHousingCRUD(db *gorm.DB) *repositoryWarehousingCRUD {
	return repositoryWarehousingCRUD{db}
}

func (r *repositoryWarehousingCRUD) Save(models.WareHousing warehousing) (models.WareHousing,error) {
	var err Error
	done:=make(chan bool)
	 go func (ch chan<-bool) {
		 err=r.db.Create(&warehousing).Error
		 if err!=nil {
			 ch<-false
			 return 
		 }
		 ch<-true
	 }(done)
	 if channels.Ok(done) {
		 return warehousing,nil
	 }
	 return models.WareHousing{},err

}