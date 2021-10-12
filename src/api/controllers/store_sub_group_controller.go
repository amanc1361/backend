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

func CreateStoreSubGroup(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	storesubgroup := models.StoreSubGroup{}

	err = json.Unmarshal(body, &storesubgroup)
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

	repo := crud.NewRepositoryStoreSubGroupCRUD(db)

	func(storesubgroupRepository repository.StoreSubGroupRepository) {

		storesubgroup, err = storesubgroupRepository.Save(storesubgroup)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, storesubgroup.ID))
		responses.JSON(w, http.StatusOK, storesubgroup.PublicStoreSubGroup())
	}(repo)

}
func GetStoreSubGroups(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStoreSubGroupCRUD(db)
	storesubgroupis := models.StoreSubGroups{}
	func(storesubgroupRepository repository.StoreSubGroupRepository) {

		storesubgroupis, err = storesubgroupRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storesubgroupis.PublicStoreSubGroups())
	}(repo)
}

func GetStoreSubGroup(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryStoreSubGroupCRUD(db)

	func(storesubgroupRepository repository.StoreSubGroupRepository) {

		storesubgroup, err := storesubgroupRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storesubgroup)
	}(repo)
}

func UpdateStoreSubGroup(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	storesubgroup := models.StoreSubGroup{}

	err = json.Unmarshal(body, &storesubgroup)
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

	repo := crud.NewRepositoryStoreSubGroupCRUD(db)

	func(storesubgroupRepository repository.StoreSubGroupRepository) {

		updatestoresubgroup, err := storesubgroupRepository.Update(storesubgroup)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, updatestoresubgroup)
	}(repo)

}

func DeleteStoreSubGroup(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryStoreSubGroupCRUD(db)

	func(storesubgroupRepository repository.StoreSubGroupRepository) {

		storesubgroup, err := storesubgroupRepository.Delete(int32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storesubgroup)
	}(repo)
}
