package controllers

import (
	"back-account/src/api/database"
	"back-account/src/api/models"
	"back-account/src/api/repository"
	"back-account/src/api/repository/crud"
	"back-account/src/api/responses"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetBilans(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	modeltype, err := strconv.ParseInt(v.Get("modeltype"), 10, 32)
	reportbase, err := strconv.ParseInt(v.Get("reportbase"), 10, 32)
	parentid, err := strconv.ParseInt(v.Get("parentid"), 10, 32)
	yearid, err := strconv.ParseInt(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseInt(v.Get("companyid"), 10, 32)
	docfrom, err := strconv.ParseInt(v.Get("docfrom"), 10, 32)
	docto, err := strconv.ParseInt(v.Get("docto"), 10, 32)
	solarfrom := v.Get("solarfrom")
	solarto := v.Get("solarto")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("پارامترهای ورودی نادرست است"))
		return
	}

	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBilan(db)

	func(bilanRepository repository.BilanRepository) {

		bilans, err := bilanRepository.FindAll(int(reportbase), int(modeltype), int(companyid), int(yearid), int(parentid), int(docfrom), int(docto), solarfrom, solarto)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, bilans)
	}(repo)
}

func GetGroupBilans(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseInt(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseInt(v.Get("companyid"), 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("پارامترهای ورودی نادرست است"))
		return
	}

	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBilan(db)

	func(bilanRepository repository.BilanRepository) {

		bilans, err := bilanRepository.GetGroupTaraz(int(companyid), int(yearid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, bilans)
	}(repo)
}

func GetCsv(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBilan(db)

	func(bilanRepository repository.BilanRepository) {

		bilans, err := bilanRepository.Getcsv(5, 14, "", "")
		if err != nil {

			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		b := new(bytes.Buffer)

		w1 := csv.NewWriter(b)
		header := []string{"شماره سند",
			"تاریخ سند",
			"کد کل",
			"نام کل",
			"کد معین",
			"نام معین",
			"کد تفضیل",
			"نام تفضیل",
			"بدهکار",
			"بستانکار",
		}
		bomUtf8 := []byte{0xEF, 0xBB, 0xBF}
		w1.Write([]string{string(bomUtf8[:])})
		w1.Write(header)
		for _, item := range bilans {
			var s []string
			s = append(s, fmt.Sprint(item.DocumentNumber),
				item.SolarDate,
				item.LedgerCode,
				item.LedgerName,
				item.SubLedgerCode,
				item.SubLedgerName,
				item.DetailedCode,
				item.DetailedName,
				fmt.Sprint(item.Debtor),
				fmt.Sprint(item.Creditor),
			)

			w1.Write(s)
		}

		w.Header().Set("Content-Type", "text/csv") // setting the content type header to text/csv
		w.Header().Set("Content-Disposition", "attachment;filename=TheCSVFileName.csv")
		w.Write(b.Bytes())

	}(repo)

}

func GetProfit(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseInt(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseInt(v.Get("companyid"), 10, 32)
	istemp, err := strconv.ParseInt(v.Get("istemp"), 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("پارامترهای ورودی نادرست است"))
		return
	}

	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBilan(db)

	func(bilanRepository repository.BilanRepository) {

		profits, err := bilanRepository.GetProfit(int(companyid), int(yearid), int(istemp))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, profits)
	}(repo)
}

func GetTaraz(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	tabletype, err := strconv.ParseInt(v.Get("tabletype"), 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("نوع حساب را مشخص کنید"))
		return
	}
	yearid, err := strconv.ParseInt(v.Get("yearid"), 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("سال مالی را وارد نمایید"))
		return
	}
	companyid, err := strconv.ParseInt(v.Get("companyid"), 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New(" نام شرکت را مشخص کنید"))
		return
	}
	firstdoc, err := strconv.ParseBool(v.Get("firstdoc"))
	if err != nil {
		firstdoc = false
	}

	docfrom, err := strconv.ParseInt(v.Get("docfrom"), 10, 32)
	if err != nil {
		docfrom = 0
	}
	docto, err := strconv.ParseInt(v.Get("docto"), 10, 32)
	if err != nil {
		docto = 0
	}
	solarfrom := v.Get("solarfrom")
	solarto := v.Get("solarto")
	parentid, err := strconv.ParseInt(v.Get("parentid"), 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New(" نام حساب مرجع را مشخص کنید"))
		return
	}
	reporttype, err := strconv.ParseInt(v.Get("reporttype"), 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New(" نام حساب مرجع را مشخص کنید"))
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("پارامترهای ورودی نادرست است"))
		return
	}

	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBilan(db)

	func(bilanRepository repository.BilanRepository) {

		taraz, err := bilanRepository.FindTaraz(bool(firstdoc), int(tabletype), int(yearid), int(companyid), int(docfrom), int(docto), solarfrom, solarto, int(parentid), int(reporttype))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, taraz)
	}(repo)
}

func GetDocByDetaildId(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	yearid, err := strconv.ParseInt(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseInt(v.Get("companyid"), 10, 32)
	itemid, err := strconv.ParseInt(v.Get("itemid"), 10, 32)
	subledgerid, err := strconv.ParseInt(v.Get("subledgerid"), 10, 32)
	tmodel, err := strconv.ParseInt(v.Get("tmodel"), 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("پارامترهای ورودی نادرست است"))
		return
	}

	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBilan(db)

	func(bilanRepository repository.BilanRepository) {

		docs, err := bilanRepository.FindDetaildsById(int(tmodel), int(yearid), int(companyid), int(itemid), int(subledgerid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, docs)
	}(repo)
}

func GetBilanBySearch(w http.ResponseWriter, r *http.Request) {

	search := models.Searchs{}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var v = r.URL.Query()

	modeltype, err := strconv.ParseInt(v.Get("modeltype"), 10, 32)
	yearid, err := strconv.ParseInt(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseInt(v.Get("companyid"), 10, 32)
	parentid, err := strconv.ParseInt(v.Get("parentid"), 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("پارامترهای ورودی نادرست است"))
	}
	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = json.Unmarshal(body, &search)
	if err != nil {
		search.Serach = ""
	}

	repo := crud.NewRepositoryBilan(db)

	func(bilanRepository repository.BilanRepository) {

		bilans, err := bilanRepository.FindBySearch(int(modeltype), int(yearid), int(companyid), search.Serach, int(parentid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, bilans)
	}(repo)
}

func GetTaraNameh(w http.ResponseWriter, r *http.Request) {

	var v = r.URL.Query()

	reportbase, err := strconv.ParseInt(v.Get("reportbase"), 10, 32)
	yearid, err := strconv.ParseInt(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseInt(v.Get("companyid"), 10, 32)
	solarto := v.Get("solarto")

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("پارامترهای ورودی نادرست است"))
	}
	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBilan(db)

	func(bilanRepository repository.BilanRepository) {

		bilans, err := bilanRepository.GetTaraNameh(int(companyid), int(yearid), int(reportbase), solarto)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, bilans)
	}(repo)
}

func GetProfitYears(w http.ResponseWriter, r *http.Request) {

	var v = r.URL.Query()

	reportbase, err := strconv.ParseInt(v.Get("reportbase"), 10, 32)
	yearid, err := strconv.ParseInt(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseInt(v.Get("companyid"), 10, 32)
	solarto := v.Get("solarto")
	solarfrom := v.Get("solarfrom")

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("پارامترهای ورودی نادرست است"))
	}
	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBilan(db)

	func(bilanRepository repository.BilanRepository) {

		bilans, err := bilanRepository.GetProfitYear(int(companyid), int(yearid), int(reportbase), solarfrom, solarto)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, bilans)
	}(repo)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	var v = r.URL.Query()

	reportbase, err := strconv.ParseInt(v.Get("reportbase"), 10, 32)
	yearid, err := strconv.ParseInt(v.Get("yearid"), 10, 32)
	companyid, err := strconv.ParseInt(v.Get("companyid"), 10, 32)
	parentid, err := strconv.ParseInt(v.Get("parentid"), 10, 32)
	solarto := v.Get("solarto")
	solarfrom := v.Get("solarfrom")
	docfrom, err := strconv.ParseInt(v.Get("docfrom"), 10, 32)
	docto, err := strconv.ParseInt(v.Get("docto"), 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("پارامترهای ورودی نادرست است"))
	}
	db, err := database.Connect()
	sqldb, err := db.DB()
	defer sqldb.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBilan(db)

	func(bilanRepository repository.BilanRepository) {

		articles, err := bilanRepository.GetArticles(int(companyid), int(yearid), int(reportbase), int(parentid), solarfrom, solarto, int(docfrom), int(docto))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, articles)
	}(repo)
}
