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

func CreateStorePerson(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	storeperson := models.StorePerson{}

	err = json.Unmarshal(body, &storeperson)
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

	repo := crud.NewRepositoryStorePersonCRUD(db)

	func(storepersonRepository repository.StorePerson) {

		storeperson, err = storepersonRepository.Save(storeperson)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, storeperson.ID))
		responses.JSON(w, http.StatusOK, storeperson.PublicStorePerson())
	}(repo)

}
func GetStorePersons(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	search := v.Get("search")
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStorePersonCRUD(db)
	storepersonis := models.StorePersonis{}
	func(storepersonRepository repository.StorePerson) {

		storepersonis, err = storepersonRepository.FindAll(int(companyid), search)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storepersonis.PublicStorePersonis())
	}(repo)
}

func GetStorePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStorePersonCRUD(db)

	func(storepersonRepository repository.StorePerson) {

		storeperson, err := storepersonRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storeperson)
	}(repo)
}

func UpdateStorePerson(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	storeperson := models.StorePerson{}

	err = json.Unmarshal(body, &storeperson)
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

	repo := crud.NewRepositoryStorePersonCRUD(db)

	func(storepersonRepository repository.StorePerson) {

		updatestoreperson, err := storepersonRepository.Update(storeperson)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, updatestoreperson)
	}(repo)

}

func DeleteStorePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStorePersonCRUD(db)

	func(storepersonRepository repository.StorePerson) {

		storeperson, err := storepersonRepository.Delete(int32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storeperson)
	}(repo)
}
