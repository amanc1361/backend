package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var storeGroupRoutes = []Route{
	Route{
		Uri:     "/storegroup/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storegroups",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreGroups,
		IsAuth:  true,
	},

	Route{
		Uri:     "/storegroup",
		Method:  http.MethodPost,
		Handler: controllers.CreateStoreGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storegroup/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateStoreGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storegroup/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteStoreGroup,
		IsAuth:  true,
	},
}
