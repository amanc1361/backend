package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var storeRoutes = []Route{
	Route{
		Uri:     "/store/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetStore,
		IsAuth:  true,
	},
	Route{
		Uri:     "/stores",
		Method:  http.MethodPost,
		Handler: controllers.GetStores,
		IsAuth:  true,
	},

	Route{
		Uri:     "/store",
		Method:  http.MethodPost,
		Handler: controllers.CreateStore,
		IsAuth:  true,
	},
	Route{
		Uri:     "/store/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateStore,
		IsAuth:  true,
	},
	Route{
		Uri:     "/store/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteStore,
		IsAuth:  true,
	},
}
