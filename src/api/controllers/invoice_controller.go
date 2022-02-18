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
	"strconv"
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

func GetInovices(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	invoicetype, err := strconv.ParseUint(v.Get("invoicetype"), 10, 32)
	



	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	repo := crud.NewInvocieRepository(db)

	func(invoice repository.Inovice) {

		invoices, err := invoice.GetAll( int(companyid),int(yearid),int(invoicetype))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, invoices)
	}(repo)
}




