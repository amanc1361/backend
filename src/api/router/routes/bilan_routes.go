package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var bilanRoutes = []Route{
	Route{
		Uri:     "/bilan",
		Method:  http.MethodPost,
		Handler: controllers.GetBilans,
		IsAuth:  true,
	},
	Route{
		Uri:     "/bilan/group",
		Method:  http.MethodPost,
		Handler: controllers.GetGroupBilans,
		IsAuth:  true,
	},
	Route{
		Uri:     "/taraz",
		Method:  http.MethodPost,
		Handler: controllers.GetTaraz,
		IsAuth:  true,
	},
	Route{
		Uri:     "/bilan/profit",
		Method:  http.MethodPost,
		Handler: controllers.GetProfit,
		IsAuth:  true,
	},
	Route{
		Uri:     "/bilan/search",
		Method:  http.MethodPost,
		Handler: controllers.GetBilanBySearch,
		IsAuth:  true,
	},
	Route{
		Uri:     "/bilan/detaild",
		Method:  http.MethodPost,
		Handler: controllers.GetDocByDetaildId,
		IsAuth:  true,
	},
	Route{
		Uri:     "/bilan/taraznameh",
		Method:  http.MethodPost,
		Handler: controllers.GetTaraNameh,
		IsAuth:  true,
	},
	Route{
		Uri:     "/bilan/profityear",
		Method:  http.MethodPost,
		Handler: controllers.GetProfitYears,
		IsAuth:  true,
	},
	Route{
		Uri:     "/bilan/articles",
		Method:  http.MethodPost,
		Handler: controllers.GetArticles,
		IsAuth:  true,
	},
}
