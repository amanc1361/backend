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

func CreateStoreAction(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	storeaction := models.StoreAction{}

	err = json.Unmarshal(body, &storeaction)
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

	repo := crud.NewRepositoryStoreActionCRUD(db)

	func(storeactionRepository repository.StoreActionRepository) {

		storeaction, err = storeactionRepository.Save(storeaction)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, storeaction.ID))
		responses.JSON(w, http.StatusOK, storeaction)
	}(repo)

}
func GetStoreIdByDocumentid(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()
	documentid, err := strconv.ParseUint(v.Get("documentid"), 10, 32)
	fmt.Println("--------------------------------------")
	fmt.Println(documentid)
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStoreActionCRUD(db)

	func(storeactionRepository repository.StoreActionRepository) {

		storeid, storetypeid, err := storeactionRepository.GetStoreIdbyDocumentID(int(documentid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		fmt.Println(storeid)
		storetypeid++

		responses.JSON(w, http.StatusOK, storeid)
	}(repo)
}
func GetStoreActions(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	typeactionid, err := strconv.ParseUint(v.Get("typeaction"), 10, 32)
	storeid, err := strconv.ParseUint(v.Get("storeid"), 10, 32)
	search := v.Get("search")
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStoreActionCRUD(db)
	storeactionis := models.StoreActions{}
	func(storeactionRepository repository.StoreActionRepository) {

		storeactionis, err = storeactionRepository.FindAll(search, int(typeactionid), int(yearid), int(companyid), int(storeid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storeactionis.PublicStoreActions())
	}(repo)
}
func Getsends(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)

	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStoreActionCRUD(db)
	res := []modelsout.StoreActiondoc{}
	func(storeactionRepository repository.StoreActionRepository) {

		res, err = storeactionRepository.Getsend(int(companyid), int(yearid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, res)
	}(repo)
}

func Getrecives(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)

	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStoreActionCRUD(db)
	res := []modelsout.StoreActiondoc{}
	func(storeactionRepository repository.StoreActionRepository) {

		res, err = storeactionRepository.Getrecive(int(companyid), int(yearid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, res)
	}(repo)
}

func GetStoreAction(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryStoreActionCRUD(db)

	func(storeactionRepository repository.StoreActionRepository) {

		storeaction, err := storeactionRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storeaction)
	}(repo)
}
func GetStoreActionRows(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err1 := database.Connect()
	if err1 != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	sqlDB, err2 := db.DB()
	if err2 != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStoreActionCRUD(db)

	func(storeactionRepository repository.StoreActionRepository) {

		storeaction, err := storeactionRepository.GetStoreActionRows(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storeaction)
	}(repo)
}

func UpdateStoreAction(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	storeaction := models.StoreAction{}

	err = json.Unmarshal(body, &storeaction)
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

	repo := crud.NewRepositoryStoreActionCRUD(db)

	func(storeactionRepository repository.StoreActionRepository) {

		updatestoreaction, err := storeactionRepository.Update(storeaction)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, updatestoreaction)
	}(repo)

}

func GetCountObjectbyid(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	uid, err := strconv.ParseUint(v.Get("id"), 10, 32)
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStoreActionCRUD(db)

	func(storeactionRepository repository.StoreActionRepository) {

		count, err := storeactionRepository.GetCountObject(int(companyid), int(yearid), int(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, count)
	}(repo)
}
func GetKardexbyid(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	uid, err := strconv.ParseUint(v.Get("id"), 10, 32)
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStoreActionCRUD(db)

	func(storeactionRepository repository.StoreActionRepository) {

		kardex, err := storeactionRepository.GetKardex(int(companyid), int(yearid), int(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, kardex)
	}(repo)
}
func GetRemObjectBtStoreid(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	uid, err := strconv.ParseUint(v.Get("id"), 10, 32)
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStoreActionCRUD(db)

	func(storeactionRepository repository.StoreActionRepository) {

		remobjects, err := storeactionRepository.GetRemObject(int(companyid), int(yearid), int(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, remobjects)
	}(repo)
}

func GetPriceObjectById(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	uid, err := strconv.ParseUint(v.Get("objectid"), 10, 32)
	solar := v.Get("solardate")
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryStoreActionCRUD(db)

	func(storeactionRepository repository.StoreActionRepository) {

		price, err := storeactionRepository.GetPriceObject(int(companyid), int(yearid), int(uid), solar)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, price)
	}(repo)
}
func DeleteStoreAction(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryStoreActionCRUD(db)

	func(storeactionRepository repository.StoreActionRepository) {

		storeaction, err := storeactionRepository.Delete(int32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storeaction)
	}(repo)
}

func GetStoreActionsAll(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	reporttype, err := strconv.ParseUint(v.Get("reporttype"), 10, 32)
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	repo := crud.NewRepositoryStoreActionCRUD(db)
	func(storeactionRepository repository.StoreActionRepository) {
		storeactions, err := storeactionRepository.GetAll(int(companyid), int(yearid), int(reporttype))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, storeactions)
	}(repo)
}
