package controllers

import (
	"back-account/src/api/database"
	"back-account/src/api/repository"
	"back-account/src/api/repository/crud"
	"back-account/src/api/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCompanyTypes(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryCompanyType(db)

	func(companyTypeRepository repository.CompanyTypesRepository) {

		companytypes, err := companyTypeRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusCreated, companytypes)
	}(repo)
}

func GetCompanyType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryCompanyType(db)

	func(companyTypeRepository repository.CompanyTypesRepository) {

		companyType, err := companyTypeRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusCreated, companyType)
	}(repo)
}
