package controllers

import (
	"back-account/src/api/database"
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
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

func GetSubLedgers(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositorySubLedgerCRUD(db)

	func(subledgerRepository repository.SubLedgerRepository) {

		subledgers, err := subledgerRepository.FindAll(int(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, subledgers)
	}(repo)
}
func CreateSubLedger(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	subledger := models.SubLedger{}
	outsubledger := modelsout.SubledgerOut{}

	err = json.Unmarshal(body, &subledger)
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

	repo := crud.NewRepositorySubLedgerCRUD(db)

	func(subledgerRepository repository.SubLedgerRepository) {

		outsubledger, err = subledgerRepository.Save(subledger)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, outsubledger.ID))
		responses.JSON(w, http.StatusOK, outsubledger)
	}(repo)
}
func UpdateSubLedger(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	subledger := models.SubLedger{}

	err = json.Unmarshal(body, &subledger)
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

	repo := crud.NewRepositorySubLedgerCRUD(db)

	func(subledgerRepository repository.SubLedgerRepository) {

		updatecount, err := subledgerRepository.Update(subledger)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, updatecount)
	}(repo)
}
func DeleteSubLedger(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositorySubLedgerCRUD(db)

	func(subledgerRepository repository.SubLedgerRepository) {

		subledgers, err := subledgerRepository.Delete(int32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, subledgers)
	}(repo)
}
func GetSubLedger(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositorySubLedgerCRUD(db)

	func(subledgerRepository repository.SubLedgerRepository) {

		subledger, err := subledgerRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, subledger)
	}(repo)
}
func GetSubLedgerByCode(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["code"], 10, 64)
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

	repo := crud.NewRepositorySubLedgerCRUD(db)

	func(subledgerRepository repository.SubLedgerRepository) {

		subledgers, err := subledgerRepository.FindByCode(uint(uid), int(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, subledgers)
	}(repo)
}

func GetLastSubLeadgercode(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
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

	repo := crud.NewRepositorySubLedgerCRUD(db)

	func(subledgerRepository repository.SubLedgerRepository) {

		code, err := subledgerRepository.GetLastCode(uint(id), int(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, code)
	}(repo)
}

func GetByLedgerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
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

	repo := crud.NewRepositorySubLedgerCRUD(db)

	func(subledgerRepository repository.SubLedgerRepository) {

		subledgers, err := subledgerRepository.GetByLedgerID(uint(id))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, subledgers)
	}(repo)
}
