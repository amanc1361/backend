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

func CreateStorePersonRecive(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	storepersonrecive := models.StorePersonRecive{}

	err = json.Unmarshal(body, &storepersonrecive)
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

	repo := crud.NewRepositoryStorePersonReciveCRUD(db)

	func(storepersonreciveRepository repository.StorePersonRecive) {

		storepersonrecive, err = storepersonreciveRepository.Save(storepersonrecive)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, storepersonrecive.ID))
		responses.JSON(w, http.StatusOK, storepersonrecive.PublicStorePersonRecive())
	}(repo)

}
func GetStorePersonRecives(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryStorePersonReciveCRUD(db)
	storepersonreciveis := models.StorePersonReciveis{}
	func(storepersonreciveRepository repository.StorePersonRecive) {

		storepersonreciveis, err = storepersonreciveRepository.FindAll(int(companyid), search)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storepersonreciveis.PublicStorePersonReciveis())
	}(repo)
}

func GetStorePersonRecive(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryStorePersonReciveCRUD(db)

	func(storepersonreciveRepository repository.StorePersonRecive) {

		storepersonrecive, err := storepersonreciveRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storepersonrecive)
	}(repo)
}

func UpdateStorePersonRecive(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	storepersonrecive := models.StorePersonRecive{}

	err = json.Unmarshal(body, &storepersonrecive)
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

	repo := crud.NewRepositoryStorePersonReciveCRUD(db)

	func(storepersonreciveRepository repository.StorePersonRecive) {

		updatestorepersonrecive, err := storepersonreciveRepository.Update(storepersonrecive)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, updatestorepersonrecive)
	}(repo)

}

func DeleteStorePersonRecive(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryStorePersonReciveCRUD(db)

	func(storepersonreciveRepository repository.StorePersonRecive) {

		storepersonrecive, err := storepersonreciveRepository.Delete(int32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storepersonrecive)
	}(repo)
}
