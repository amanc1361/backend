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

func CreatePeople(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	person := models.Person{}

	err = json.Unmarshal(body, &person)
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

	repo := crud.NewRepositoryPeopleCRUD(db)

	func(storepersonRepository repository.PeopleRepository) {

		person, err = storepersonRepository.Save(person)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, person.ID))
		responses.JSON(w, http.StatusOK, person)
	}(repo)

}
func GetPeoples(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	search := v.Get("search")
	db, err := database.Connect()
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPeopleCRUD(db)

	func(storepersonRepository repository.PeopleRepository) {

		storepersonis, err := storepersonRepository.FindAll(int(companyid), search)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storepersonis)
	}(repo)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryPeopleCRUD(db)

	func(storepersonRepository repository.PeopleRepository) {

		storeperson, err := storepersonRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storeperson)
	}(repo)
}

func UpdatePeople(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	storeperson := models.Person{}

	err = json.Unmarshal(body, &storeperson)
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

	repo := crud.NewRepositoryPeopleCRUD(db)

	func(storepersonRepository repository.PeopleRepository) {

		updatestoreperson, err := storepersonRepository.Update(storeperson)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, updatestoreperson)
	}(repo)

}

func GetRemPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("00000000000000000000000000000000000000000000000")
	var v = r.URL.Query()
	companyid, err := strconv.ParseUint(v.Get("companyid"), 10, 32)
	fmt.Println(companyid)
	yearid, err := strconv.ParseUint(v.Get("yearid"), 10, 32)
	fmt.Println(yearid)
	detailedid, err := strconv.ParseUint(v.Get("detailedid"), 10, 32)
	fmt.Println(detailedid)
	solardate:=v.Get("solardate")
  fmt.Println(solardate)
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

	repo := crud.NewRepositoryPeopleCRUD(db)

	func(storepersonRepository repository.PeopleRepository) {

		rem, err := storepersonRepository.GetRemPerson(int(companyid),int(yearid),int(detailedid),solardate)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}



		responses.JSON(w, http.StatusOK, rem)
	}(repo)
}



func DeletePeople(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()
	peopleid, err := strconv.ParseUint(v.Get("id"), 10, 32)
	detailedid, err := strconv.ParseUint(v.Get("detailedid"), 10, 32)

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

	repo := crud.NewRepositoryPeopleCRUD(db)

	func(storepersonRepository repository.PeopleRepository) {

		storeperson, err := storepersonRepository.Delete(int32(peopleid), int32(detailedid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, storeperson)
	}(repo)
}
