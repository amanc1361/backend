package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var storeRoutes = []Route{
	{
		Uri:     "/store/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetStore,
		IsAuth:  true,
	},
	{
		Uri:     "/stores",
		Method:  http.MethodPost,
		Handler: controllers.GetStores,
		IsAuth:  true,
	},

	{
		Uri:     "/store",
		Method:  http.MethodPost,
		Handler: controllers.CreateStore,
		IsAuth:  true,
	},
	{
		Uri:     "/store/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateStore,
		IsAuth:  true,
	},
	{
		Uri:     "/store/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteStore,
		IsAuth:  true,
	},
	{
		Uri:     "/storerem",
		Method:  http.MethodGet,
		Handler: controllers.GetStorerem,
		IsAuth:  true,
	},
	{
		Uri:     "/storeremobject",
		Method:  http.MethodGet,
		Handler: controllers.GetStorewithremobject,
		IsAuth:  true,
	},
	{
		Uri:     "/remstoreobjectbystoreid",
		Method:  http.MethodGet,
		Handler: controllers.GetRemStoreObjectByStoreId,
		IsAuth:  true,
	},


}
