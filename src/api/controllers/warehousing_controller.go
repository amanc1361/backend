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

func GetWareHousings(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()
	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
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
	repo:=crud.NewRepositoryWareHousingCRUD(db)
	func(wareHousing repository.WareHousingRepository) {

		warehusings, err := wareHousing.FindAll(int(companyid),int(yearid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, warehusings)
	}(repo)
}
func SaveWareHouse(w http.ResponseWriter,r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	warehouing := models.WareHousing{}
	err = json.Unmarshal(body, &warehouing)
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
	repo:=crud.NewRepositoryWareHousingCRUD(db)
	func(wareHousing repository.WareHousingRepository) {
		warehouing, err = wareHousing.Save(warehouing)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, wareHousing)
	}(repo)
}

func UpdateWareHouse(w http.ResponseWriter,r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	warehouing := models.WareHousing{}
	err = json.Unmarshal(body, &warehouing)
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
	repo:=crud.NewRepositoryWareHousingCRUD(db)
	func(wareHousing repository.WareHousingRepository) {
		warehouingid, err := wareHousing.Update(warehouing)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, warehouingid)
	}(repo)
}

func GetWareHouseById(w http.ResponseWriter,r *http.Request) {
	var v = r.URL.Query()
	id, err := strconv.ParseUint(v.Get("id"), 10, 32)
	
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
	repo:=crud.NewRepositoryWareHousingCRUD(db)
	func(wareHousing repository.WareHousingRepository) {

		warehusing, err := wareHousing.GetWareHousingById(int(id))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, warehusing)
	}(repo)
}