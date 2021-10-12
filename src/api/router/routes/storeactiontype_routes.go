package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var storeactiontypeRoutes = []Route{
	Route{
		Uri:     "/storeactiontype/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreActionType,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeactiontypes",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreActionTypes,
		IsAuth:  true,
	},

	Route{
		Uri:     "/storeactiontype",
		Method:  http.MethodPost,
		Handler: controllers.CreateStoreActionType,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeactiontype/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateStoreActionType,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeactiontype/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteStoreActionType,
		IsAuth:  true,
	},
}
