package crud

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
	"back-account/src/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryBilan struct {
	db *gorm.DB
}

func NewRepositoryBilan(db *gorm.DB) *repositoryBilan {
	return &repositoryBilan{db}
}

func (r *repositoryBilan) FindDetaildsById(tmodel int, yearid int, companyid int, itemid int, subledgerid int) ([]modelsout.Doc, error) {
	var err error
	var result *gorm.DB
	docs := []modelsout.Doc{}
	done := make(chan bool)

	go func(ch chan<- bool) {

		result = r.db.Raw("call getdocumentsbydetailedid(?,?,?,?,?);", tmodel, yearid, companyid, itemid, subledgerid).Scan(&docs)

		if result.Error != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return docs, nil
	}

	return nil, err
}
func (r *repositoryBilan) FindAll(reportbase int, typemodel int, companyid int, yearid int, parentid int, docfrom int, docto int, solarfrom string, solarto string) ([]models.Bilan, error) {

	var err error
	var result *gorm.DB
	bilans := []models.Bilan{}
	done := make(chan bool)

	go func(ch chan<- bool) {

		result = r.db.Debug().Raw("call getbilans(?,?,?,?,?,?,?,?,?);", reportbase, typemodel, companyid, yearid, parentid, docfrom, docto, solarfrom, solarto).Scan(&bilans)

		if result.Error != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return bilans, nil
	}

	return nil, err

}

func (r *repositoryBilan) FindBySearch(typemodel int, yearid int, companyid int, search string, parentid int) ([]models.Bilan, error) {

	var err error
	var result *gorm.DB
	bilans := []models.Bilan{}
	done := make(chan bool)

	go func(ch chan<- bool) {

		switch typemodel {
		case 1:
			result = r.db.Raw("call getgroupbilan(?,?);", yearid, companyid).Scan(&bilans)
		case 2:
			result = r.db.Raw("call getledgerbilan(?,?);", yearid, companyid).Scan(&bilans)
		case 3:
			result = r.db.Raw("call getsubledgerbilan(?,?);", yearid, companyid).Scan(&bilans)
		case 4:
			result = r.db.Raw("call getdetailedbilan(?,?);", yearid, companyid).Scan(&bilans)
		case 5:
			result = r.db.Raw("call getdetailedbilanbyledgerid(?,?,?);", yearid, companyid, parentid).Scan(&bilans)
		}

		if result.Error != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return bilans, nil
	}

	return nil, err

}

func (r *repositoryBilan) FindTaraz(firstdoc bool, tabletype int, yearid int, companyid int, docfrom int, docto int, solarfrom string, solarto string, parentid int, reporttype int) ([]modelsout.Taraz, error) {

	var err error
	var result *gorm.DB
	taraz := []modelsout.Taraz{}
	done := make(chan bool)

	go func(ch chan<- bool) {

		result = r.db.Raw("call getBilan(?,?,?,?,?,?,?,?,?,?);", firstdoc, tabletype, yearid, companyid, docfrom, docto, solarfrom, solarto, parentid, reporttype).Scan(&taraz)

		if result.Error != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return taraz, nil
	}

	return nil, err

}

func (r *repositoryBilan) GetGroupTaraz(companyid int, yearid int) ([]models.GroupTaraz, error) {

	var err error
	var result *gorm.DB
	grouptaraz := []models.GroupTaraz{}
	done := make(chan bool)

	go func(ch chan<- bool) {

		result = r.db.Raw("call getgroupingbilan(?,?);", companyid, yearid).Scan(&grouptaraz)

		if result.Error != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return grouptaraz, nil
	}

	return nil, err

}

func (r *repositoryBilan) GetProfit(companyid int, yearid int, istemp int) ([]modelsout.Profit, error) {

	var err error
	var result *gorm.DB
	profits := []modelsout.Profit{}
	done := make(chan bool)

	go func(ch chan<- bool) {

		result = r.db.Raw("call getprofit(?,?,?);", companyid, yearid, istemp).Scan(&profits)

		if result.Error != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return profits, nil
	}

	return nil, err

}
func (r *repositoryBilan) GetTaraNameh(companyid int, yearid int, reportbase int, soalrto string) ([]models.Tarznameh, error) {

	var err error
	var result *gorm.DB
	tarznameh := []models.Tarznameh{}
	done := make(chan bool)

	go func(ch chan<- bool) {

		result = r.db.Raw("call gettarznameh(?,?,?,?);", companyid, yearid, reportbase, soalrto).Scan(&tarznameh)

		if result.Error != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return tarznameh, nil
	}

	return nil, err

}
func (r *repositoryBilan) GetProfitYear(companyid int, yearid int, reportbase int, solarfrom string, soalrto string) ([]models.Tarznameh, error) {

	var err error
	var result *gorm.DB
	tarznameh := []models.Tarznameh{}
	done := make(chan bool)

	go func(ch chan<- bool) {

		result = r.db.Raw("call getprofityears(?,?,?,?,?);", companyid, yearid, reportbase, solarfrom, soalrto).Scan(&tarznameh)

		if result.Error != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return tarznameh, nil
	}

	return nil, err

}

func (r *repositoryBilan) GetArticles(companyid int, yearid int, reportbase int, parentid int, solarfrom string, solarto string, docfrom int, docto int) ([]models.Article, error) {

	var err error
	var result *gorm.DB
	articles := []models.Article{}
	done := make(chan bool)

	go func(ch chan<- bool) {

		result = r.db.Debug().Raw("call getdocumentrows(?,?,?,?,?,?,?,?)", companyid, yearid, reportbase, parentid, solarfrom, solarto, docfrom, docto).Scan(&articles)

		if result.Error != nil {
			ch <- false
			return
		}
		ch <- true

	}(done)

	if channels.Ok(done) {
		return articles, nil
	}

	return nil, err
}
func (r *repositoryBilan) Getcsv(companyid int, yearid int, solarfrom string, solarto string) ([]models.DocumnetCsv, error) {
	var err error
	csv := []models.DocumnetCsv{}

	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Table("document_rows").
			Joins("join documents on document_rows.document_id=documents.id ").
			Joins(" join ledgers on document_rows.ledger_id=ledgers.id").
			Joins(" join sub_ledgers on document_rows.sub_ledger_id=sub_ledgers.id").
			Joins(" left join detaileds on document_rows.detailed_id=detaileds.id").
			Select(` documents.document_number as DocumentNumber,
					documents.solar_date as SolarDate,
					ledgers.code as LedgerCode,
					ledgers.name as LedgerName,
					sub_ledgers.code as SubLedgerCode,
					sub_ledgers.name as SubLedgerName,
					detaileds.code as DetailedCode,
					detaileds.name as DetailedName,
					document_rows.debtor,
					document_rows.creditor `).
			Where("documents.company_id=? and documents.year_id=? and documents.deleted_at is null", companyid, yearid).
			Find(&csv).Error

		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)

	if channels.Ok(done) {
		return csv, nil

	}
	return nil, err

}
