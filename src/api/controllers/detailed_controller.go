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

func GetDetaileds(w http.ResponseWriter, r *http.Request) {
	var err error
	var v = r.URL.Query()

	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)

	search := models.Searchs{}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &search)
	if err != nil {
		search.Serach = ""
	}
	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryDetailedCRUD(db)

	func(detailedRepository repository.DetailedRepository) {

		detaileds, pagecount, err := detailedRepository.FindAll(search.Serach, int(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSONWITHPAGENUMBER(w, http.StatusOK, detaileds, pagecount)
	}(repo)
}
func CreateDetailed(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	detailed := models.Detailed{}

	err = json.Unmarshal(body, &detailed)
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

	repo := crud.NewRepositoryDetailedCRUD(db)

	func(detailedRepository repository.DetailedRepository) {

		detailed, err = detailedRepository.Save(detailed)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, detailed.ID))
		responses.JSON(w, http.StatusOK, detailed)
	}(repo)
}
func UpdateDetailed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	detailed := models.Detailed{}

	err = json.Unmarshal(body, &detailed)
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

	repo := crud.NewRepositoryDetailedCRUD(db)

	func(detailedRepository repository.DetailedRepository) {

		updatecount, err := detailedRepository.Update(int32(uid), detailed)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, updatecount)
	}(repo)
}
func DeleteDetailed(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryDetailedCRUD(db)

	func(detailedRepository repository.DetailedRepository) {

		detaileds, err := detailedRepository.Delete(int32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, detaileds)
	}(repo)
}
func GetDetailed(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryDetailedCRUD(db)

	func(detailedRepository repository.DetailedRepository) {

		detaileds, err := detailedRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, detaileds)
	}(repo)
}
func GetDetailedByCode(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["code"], 10, 64)
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

	repo := crud.NewRepositoryDetailedCRUD(db)

	func(detailedRepository repository.DetailedRepository) {

		detaileds, err := detailedRepository.FindByCode(uint(uid), int(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, detaileds)
	}(repo)
}

func GetLastDetailedcode(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryDetailedCRUD(db)

	func(detailedRepository repository.DetailedRepository) {

		code, err := detailedRepository.GetLastCode(int(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, code)
	}(repo)
}

func GetFlowDetailed(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	detailedid, err := strconv.ParseUint(v.Get("detailedid"), 10, 32)
	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryDetailedCRUD(db)

	func(detailedRepository repository.DetailedRepository) {

		detailedflows, err := detailedRepository.GetFlow(int(yearid), int(detailedid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, detailedflows)
	}(repo)
}
