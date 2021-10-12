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

func CreateStoreChildSubGroup(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	storechildsubgroup := models.StoreChildSubGroup{}

	err = json.Unmarshal(body, &storechildsubgroup)
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

	repo := crud.NewRepositoryStoreChildSubGroupCRUD(db)

	func(storechildsubgroupRepository repository.StoreChildSubGroupRepository) {

		storechildsubgroup, err = storechildsubgroupRepository.Save(storechildsubgroup)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, storechildsubgroup.ID))
		responses.JSON(w, http.StatusOK, storechildsubgroup.PublicStoreChildSubGroup())
	}(repo)

}
func GetStoreChildSubGroups(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStoreChildSubGroupCRUD(db)
	storechildsubgroupis := models.StoreChildSubGroups{}
	func(storechildsubgroupRepository repository.StoreChildSubGroupRepository) {

		storechildsubgroupis, err = storechildsubgroupRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storechildsubgroupis.PublicStoreChildSubGroups())
	}(repo)
}

func GetStoreChildSubGroup(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryStoreChildSubGroupCRUD(db)

	func(storechildsubgroupRepository repository.StoreChildSubGroupRepository) {

		storechildsubgroup, err := storechildsubgroupRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storechildsubgroup)
	}(repo)
}

func UpdateStoreChildSubGroup(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	storechildsubgroup := models.StoreChildSubGroup{}

	err = json.Unmarshal(body, &storechildsubgroup)
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

	repo := crud.NewRepositoryStoreChildSubGroupCRUD(db)

	func(storechildsubgroupRepository repository.StoreChildSubGroupRepository) {

		updatestorechildsubgroup, err := storechildsubgroupRepository.Update(storechildsubgroup)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, updatestorechildsubgroup)
	}(repo)

}

func DeleteStoreChildSubGroup(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryStoreChildSubGroupCRUD(db)

	func(storechildsubgroupRepository repository.StoreChildSubGroupRepository) {

		storechildsubgroup, err := storechildsubgroupRepository.Delete(int32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storechildsubgroup)
	}(repo)
}
