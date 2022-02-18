package controllers

import (
	"back-account/src/api/database"
	"back-account/src/api/models"
	"back-account/src/api/repository"
	"back-account/src/api/repository/crud"
	"back-account/src/api/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateInvocie(w http.ResponseWriter, r *http.Request) {
	
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	invoice := models.Invocie{}

	err = json.Unmarshal(body, &invoice)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	sqlDB, err := db.DB()

	defer sqlDB.Close()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewInvocieRepository(db)

	func(invoiceRepository repository.Inovice) {
		 invoice,err=invoiceRepository.Save(invoice)
		 
		 if err != nil {
			 responses.ERROR(w, http.StatusUnprocessableEntity, err)
			 return
		 }
 
		 responses.JSON(w, http.StatusCreated, invoice)

	}(repo)


}
