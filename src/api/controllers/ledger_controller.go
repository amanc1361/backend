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

func GetLedgers(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()
	var search string
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)

	search = v.Get("search")
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryLedgerCRUD(db)

	func(ledgerRepository repository.LedgerRepository) {

		ledgers, err := ledgerRepository.FindAll(int(companyid), search)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, ledgers)
	}(repo)
}
func CreateLedger(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	ledger := models.Ledger{}

	err = json.Unmarshal(body, &ledger)
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

	repo := crud.NewRepositoryLedgerCRUD(db)

	func(ledgerRepository repository.LedgerRepository) {

		cledger, err := ledgerRepository.Save(ledger)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, cledger.ID))
		responses.JSON(w, http.StatusOK, cledger)
	}(repo)
}
func UpdateLedger(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	ledger := models.Ledger{}

	err = json.Unmarshal(body, &ledger)
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

	repo := crud.NewRepositoryLedgerCRUD(db)

	func(ledgerRepository repository.LedgerRepository) {

		updatecount, err := ledgerRepository.Update(ledger)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, updatecount)
	}(repo)
}
func DeleteLedger(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryLedgerCRUD(db)

	func(ledgerRepository repository.LedgerRepository) {

		ledgers, err := ledgerRepository.Delete(int32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, ledgers)
	}(repo)
}
func GetLedger(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryLedgerCRUD(db)

	func(ledgerRepository repository.LedgerRepository) {

		ledgers, err := ledgerRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, ledgers)
	}(repo)
}
func GetLedgerByCode(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryLedgerCRUD(db)

	func(ledgerRepository repository.LedgerRepository) {

		ledgers, err := ledgerRepository.FindByCode(uint(uid), int(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, ledgers)
	}(repo)
}

func GetLastLeadgercode(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryLedgerCRUD(db)

	func(ledgerRepository repository.LedgerRepository) {

		code, err := ledgerRepository.GetLastCode(uint(id), int(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, code)
	}(repo)
}

func GetByGroupID(w http.ResponseWriter, r *http.Request) {

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

	repo := crud.NewRepositoryLedgerCRUD(db)

	func(ledgerRepository repository.LedgerRepository) {

		ledgers, err := ledgerRepository.GetByGroupID(uint(id))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, ledgers)
	}(repo)
}
