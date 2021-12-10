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

func CreateDocumentMerge(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	documentmerge := models.DocumentMerge{}

	err = json.Unmarshal(body, &documentmerge)
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

	repo := crud.NewDocumentMergeRepository(db)

	func(documentmergeRepository repository.DocumentMergeRepository) {

		documentmerge, err = documentmergeRepository.Save(documentmerge)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, documentmerge)
	}(repo)
}

func GetDocumentMerges(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewDocumentMergeRepository(db)

	func(document repository.DocumentMergeRepository) {

		documents, err := document.FindAll(int(companyid), int(yearid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, documents)
	}(repo)
}

func GetLastDocumentMergecode(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewDocumentMergeRepository(db)

	func(documentRepository repository.DocumentMergeRepository) {

		code, err := documentRepository.GetLastDocumetMergeCode(int(companyid), int(yearid))
		if err != nil {
			responses.ERROR(w, http.StatusCreated, err)
			return
		}

		responses.JSON(w, http.StatusOK, code)
	}(repo)
}

func DeleteDocumentMerge(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewDocumentMergeRepository(db)

	func(document repository.DocumentMergeRepository) {

		documents, err := document.Delete(int(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, documents)
	}(repo)
}

func GetLedgerDocumentByid(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	documentmerge := models.DocumentMerge{}

	err = json.Unmarshal(body, &documentmerge)
	if err != nil {
		fmt.Println(err)
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

	repo := crud.NewDocumentMergeRepository(db)

	func(documentmergeRepository repository.DocumentMergeRepository) {

		ledgerdocuments, err := documentmergeRepository.GetLedgerDocument(documentmerge)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, ledgerdocuments)
	}(repo)
}
