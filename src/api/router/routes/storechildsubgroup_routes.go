package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var storechildsubgroupRoutes = []Route{
	Route{
		Uri:     "/storechildsubgroup/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreChildSubGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storechildsubgroup/list",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreChildSubGroups,
		IsAuth:  true,
	},

	Route{
		Uri:     "/storechildsubgroup",
		Method:  http.MethodPost,
		Handler: controllers.CreateStoreChildSubGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storechildsubgroup/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateStoreChildSubGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storechildsubgroup/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteStoreChildSubGroup,
		IsAuth:  true,
	},
}
