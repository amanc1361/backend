package controllers

import (
	"back-account/src/api/database"
	"back-account/src/api/models"
	"back-account/src/api/repository"
	"back-account/src/api/repository/crud"
	"back-account/src/api/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetYears(w http.ResponseWriter, r *http.Request) {

	var v = r.URL.Query()

	companyid, _ := strconv.ParseUint(v.Get("companyid"), 10, 32)

	db, _ := database.Connect()
	sqldb, err := db.DB()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	defer sqldb.Close()

	repo := crud.NewRepositoryYearCrud(db)

	func(yearRepository repository.YearRepository) {

		years, err := yearRepository.FindAll(int(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, years)
	}(repo)
}

func CreateYear(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	year := models.Year{}

	err = json.Unmarshal(body, &year)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryYearCrud(db)

	func(yearRepository repository.YearRepository) {

		year, err = yearRepository.Save(year)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, year.ID))
		responses.JSON(w, http.StatusOK, year)
	}(repo)

}

func UpdateYear(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	year := models.Year{}

	err = json.Unmarshal(body, &year)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	sqldb, err := db.DB()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer sqldb.Close()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryYearCrud(db)

	func(yearRepository repository.YearRepository) {

		updatecount, err := yearRepository.Update(year)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, updatecount)
	}(repo)

}

func DeleteYear(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryYearCrud(db)

	func(yearRepository repository.YearRepository) {

		users, err := yearRepository.Delete(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, users)
	}(repo)
}
func GetYear(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryYearCrud(db)

	func(yearRepository repository.YearRepository) {

		users, err := yearRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, users)
	}(repo)
}
