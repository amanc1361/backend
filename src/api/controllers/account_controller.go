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

func SaveAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	account:=models.Account{}
	err = json.Unmarshal(body, &account)
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
	repo := crud.NewRepositoryAccountCRUD(db)
	func(accountRepository repository.AccountRepository) {
		account, err = accountRepository.Save(account)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, account)
	}(repo)
}
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	account:=models.Account{}
	err = json.Unmarshal(body, &account)
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
	repo := crud.NewRepositoryAccountCRUD(db)
	func(accountRepository repository.AccountRepository) {
		account, err = accountRepository.Update(account)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, account)
	}(repo)
}
func DeleteAccount(w http.ResponseWriter,r *http.Request) {
	var v = r.URL.Query()

	accountid, err := strconv.ParseUint(v.Get("accountid"), 10, 32)
		
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
		repo := crud.NewRepositoryAccountCRUD(db)
		func(accountRepository repository.AccountRepository) {
			 err = accountRepository.Delete(int(accountid))
			if err != nil {
				responses.ERROR(w, http.StatusUnprocessableEntity, err)
				return
			}
			responses.JSON(w, http.StatusOK, 1)
		}(repo)
	
}
func GetAccount(w http.ResponseWriter,r *http.Request) {
	var v = r.URL.Query()

	accountid, err := strconv.ParseUint(v.Get("accountid"), 10, 32)
		
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
		repo := crud.NewRepositoryAccountCRUD(db)
		func(accountRepository repository.AccountRepository) {
			 account,err := accountRepository.Get(int(accountid))
			if err != nil {
				responses.ERROR(w, http.StatusUnprocessableEntity, err)
				return
			}
			responses.JSON(w, http.StatusOK, account)
		}(repo)
	
}

func GetAccounts(w http.ResponseWriter,r *http.Request) {
	var v = r.URL.Query()

	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
		
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
		repo := crud.NewRepositoryAccountCRUD(db)
		func(accountRepository repository.AccountRepository) {
			 accounts,err := accountRepository.Gets(int(companyid))
			if err != nil {
				responses.ERROR(w, http.StatusUnprocessableEntity, err)
				return
			}
			responses.JSON(w, http.StatusOK, accounts)
		}(repo)
	
}