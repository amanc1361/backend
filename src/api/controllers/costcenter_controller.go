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

func GetCostCenters(w http.ResponseWriter,r *http.Request) {
	var v = r.URL.Query()
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	repo := crud.NewRepositoryCostCenter(db)
	func(costcentersRepository repository.CostCenterRepository) {
		costcenters, err := costcentersRepository.FindAll(int(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, costcenters)
	}(repo)
}  
func GetCostCenterByID(w http.ResponseWriter,r *http.Request) {
	var v = r.URL.Query()
	id, err := strconv.ParseUint(v.Get("id"), 10, 32)
	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	repo := crud.NewRepositoryCostCenter(db)
	func(costcentersRepository repository.CostCenterRepository) {
		costcenter, err := costcentersRepository.FindByID(int(id))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, costcenter)
	}(repo)
} 

func CostDelete(w http.ResponseWriter,r *http.Request) {
	var v = r.URL.Query()
	id, err := strconv.ParseUint(v.Get("id"), 10, 32)
	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	repo := crud.NewRepositoryCostCenter(db)
	func(costcentersRepository repository.CostCenterRepository) {
		costcenter, err := costcentersRepository.Delete(int(id))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, costcenter)
	}(repo)
}

func CreateCostCenter(w http.ResponseWriter,r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	costcenter := models.CostCenter{}

	err = json.Unmarshal(body, &costcenter)
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
	repo := crud.NewRepositoryCostCenter(db)
	func(costcentersRepository repository.CostCenterRepository) {
		costcenter, err := costcentersRepository.Save(costcenter)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, costcenter)
	}(repo)
}
func UpdateCostCenter(w http.ResponseWriter,r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	costcenter := models.CostCenter{}

	err = json.Unmarshal(body, &costcenter)
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
	repo := crud.NewRepositoryCostCenter(db)
	func(costcentersRepository repository.CostCenterRepository) {
		costcenter, err := costcentersRepository.Update(costcenter)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, costcenter)
	}(repo)
}
