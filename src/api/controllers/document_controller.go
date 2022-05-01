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

func CreateDocument(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	document := models.Document{}

	err = json.Unmarshal(body, &document)
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

	repo := crud.NewDocumentRepository(db)

	func(documentRepository repository.DocumentRepository) {

		document, err = documentRepository.Save(document)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, document.ID))
		responses.JSON(w, http.StatusOK, document)
	}(repo)
}
func SortDocuments(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	repo := crud.NewDocumentRepository(db)

	func(document repository.DocumentRepository) {

		result, err := document.SortDoc(int(companyid), int(yearid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, result)
	}(repo)
}

func GetDocuments(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	page, err := strconv.ParseUint(v.Get("page"), 10, 32)
	sort := v.Get("sort")

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
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	repo := crud.NewDocumentRepository(db)

	func(document repository.DocumentRepository) {

		documents, err := document.FindAll(search.Serach, uint(yearid), int(companyid), int(page), sort)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, documents)
	}(repo)
}

func GetDcoumentsByTypeDoc(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	doctype, err := strconv.ParseUint(v.Get("doctype"), 10, 32)

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
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	repo := crud.NewDocumentRepository(db)

	func(document repository.DocumentRepository) {

		documents, err := document.GetByType(int(companyid), int(yearid), int(doctype))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, documents)
	}(repo)
}

func GetDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 64)
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

	repo := crud.NewDocumentRepository(db)

	func(document repository.DocumentRepository) {

		documents, err := document.FindByID(uint(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, documents)
	}(repo)
}
func GetDocumentRowsbyDocumentid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 64)
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

	repo := crud.NewDocumentRepository(db)

	func(document repository.DocumentRepository) {

		documents, err := document.FindRowsByDcoumentid(uint(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, documents)
	}(repo)
}

func GetDocumentByCode(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
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

	repo := crud.NewDocumentRepository(db)

	func(document repository.DocumentRepository) {

		documents, err := document.FindByDocumentCode(uint(uid), int(yearid), int(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, documents)
	}(repo)
}

func GetDocumentByDescription(w http.ResponseWriter, r *http.Request) {

	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	vars := mux.Vars(r)

	uid := vars["des"]

	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewDocumentRepository(db)

	func(document repository.DocumentRepository) {

		documents, err := document.FindByDocumentDescription(uid, int(yearid), int(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, documents)
	}(repo)
}

func UpdateDocument(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	document := models.Document{}

	err = json.Unmarshal(body, &document)
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

	repo := crud.NewDocumentRepository(db)

	func(documentRepository repository.DocumentRepository) {

		documents, err := documentRepository.Update(uint(uid), int(companyid), document)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, documents)
	}(repo)
}

func GetLastDocumentcode(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 0, 64)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 0, 64)
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewDocumentRepository(db)

	func(documentRepository repository.DocumentRepository) {

		code, err := documentRepository.GetLastCode(uint(yearid), uint(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusCreated, err)
			return
		}

		responses.JSON(w, http.StatusOK, code)
	}(repo)
}

func GetFirstDate(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 0, 64)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 0, 64)
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewDocumentRepository(db)

	func(documentRepository repository.DocumentRepository) {

		solardate, err := documentRepository.GetFirstDate(int(companyid), int(yearid))
		if err != nil {
			responses.ERROR(w, http.StatusCreated, err)
			return
		}

		responses.JSON(w, http.StatusOK, solardate)
	}(repo)
}

func GetLasttDate(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 0, 64)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 0, 64)
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewDocumentRepository(db)

	func(documentRepository repository.DocumentRepository) {

		solardate, err := documentRepository.GetLastDate(int(companyid), int(yearid))
		if err != nil {
			responses.ERROR(w, http.StatusCreated, err)
			return
		}

		responses.JSON(w, http.StatusOK, solardate)
	}(repo)
}

func GetDocInfo(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseUint(v.Get("yearid"), 0, 64)
	companyid, err := strconv.ParseUint(v.Get("companyid"), 0, 64)
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewDocumentRepository(db)

	func(documentRepository repository.DocumentRepository) {

		code, err := documentRepository.GetInfo(uint(yearid), uint(companyid))
		if err != nil {
			responses.ERROR(w, http.StatusCreated, err)
			return
		}

		responses.JSON(w, http.StatusOK, code)
	}(repo)
}
func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 64)
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

	repo := crud.NewDocumentRepository(db)

	func(document repository.DocumentRepository) {

		documents, err := document.Delete(uint(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, documents)
	}(repo)
}
