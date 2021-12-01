package repository

import (
	"back-account/src/api/models"
	"back-account/src/api/modelsout"
)

type BilanRepository interface {
	FindAll(int, int,int, int,int,int, int,string,string) ([]models.Bilan, error)
	FindBySearch(int, int, int, string, int) ([]models.Bilan, error)
	FindDetaildsById(int, int, int, int, int) ([]modelsout.Doc, error)
	FindTaraz(bool, int, int, int, int, int, string, string, int, int) ([]modelsout.Taraz, error)
	GetGroupTaraz(companyid int, yearid int) ([]models.GroupTaraz, error)
	GetProfit(companyid int, yearid int, istemp int) ([]modelsout.Profit, error)
	GetTaraNameh(companyid int, yearid int, reportbase int, soalrto string) ([]models.Tarznameh, error)
	GetProfitYear(companyid int, yearid int, reportbase int, solarfrom string, soalrto string) ([]models.Tarznameh, error)
	GetArticles(companyid int, yearid int, reportbase int, parentid int, solarfrom string, solatro string, docfrom int, docto int) ([]models.Article, error)
}
