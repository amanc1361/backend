package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var detailedRoutes = []Route{
	Route{
		Uri:     "/detailed/list",
		Method:  http.MethodPost,
		Handler: controllers.GetDetaileds,
		IsAuth:  true,
	},
	Route{
		Uri:     "/detaileds/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetDetailed,
		IsAuth:  false,
	},
	Route{
		Uri:     "/detailedsbycode/{code}",
		Method:  http.MethodGet,
		Handler: controllers.GetDetailedByCode,
		IsAuth:  false,
	},
	Route{
		Uri:     "/detaileds",
		Method:  http.MethodPost,
		Handler: controllers.CreateDetailed,
		IsAuth:  true,
	},
	Route{
		Uri:     "/detaileds/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateDetailed,
		IsAuth:  true,
	},
	Route{
		Uri:     "/detaileds/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteDetailed,
		IsAuth:  true,
	},
	Route{
		Uri:     "/detaileds/lastcode",
		Method:  http.MethodPost,
		Handler: controllers.GetLastDetailedcode,
		IsAuth:  true,
	},
	Route{
		Uri:     "/detaileds/flow",
		Method:  http.MethodPost,
		Handler: controllers.GetFlowDetailed,
		IsAuth:  true,
	},
}
