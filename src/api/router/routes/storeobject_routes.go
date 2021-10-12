package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var storeObjectRoutes = []Route{
	Route{
		Uri:     "/storeobject/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreObject,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeobjects",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreObjects,
		IsAuth:  true,
	},

	Route{
		Uri:     "/customstoreobjects",
		Method:  http.MethodPost,
		Handler: controllers.GetCustomStoreObjects,
		IsAuth:  true,
	},

	Route{
		Uri:     "/storeobject",
		Method:  http.MethodPost,
		Handler: controllers.CreateStoreObject,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeobject/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateStoreObject,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeobject/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteStoreObject,
		IsAuth:  true,
	},
}
