package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var storeActionRoutes = []Route{
	Route{
		Uri:     "/storeaction/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreAction,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeactionrows/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreActionRows,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeactions",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreActions,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storesend",
		Method:  http.MethodPost,
		Handler: controllers.Getsends,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storerecive",
		Method:  http.MethodPost,
		Handler: controllers.Getrecives,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeactiongetcount",
		Method:  http.MethodPost,
		Handler: controllers.GetCountObjectbyid,
		IsAuth:  true,
	},
	Route{
		Uri:     "/priceobject",
		Method:  http.MethodPost,
		Handler: controllers.GetPriceObjectById,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeactionkardex",
		Method:  http.MethodPost,
		Handler: controllers.GetKardexbyid,
		IsAuth:  true,
	},
	Route{
		Uri:     "/getstoreidbydocumentid",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreIdByDocumentid,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeactionremobject",
		Method:  http.MethodPost,
		Handler: controllers.GetRemObjectBtStoreid,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeaction",
		Method:  http.MethodPost,
		Handler: controllers.CreateStoreAction,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeaction",
		Method:  http.MethodPut,
		Handler: controllers.UpdateStoreAction,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeactionact",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreActionsAll,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeaction/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteStoreAction,
		IsAuth:  true,
	},
}
