package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var storeSubGroupRoutes = []Route{
	Route{
		Uri:     "/storesubgroup/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreSubGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storesubgroup/list",
		Method:  http.MethodPost,
		Handler: controllers.GetStoreSubGroups,
		IsAuth:  true,
	},

	Route{
		Uri:     "/storesubgroup",
		Method:  http.MethodPost,
		Handler: controllers.CreateStoreSubGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storesubgroup/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateStoreSubGroup,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storesubgroup/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteStoreSubGroup,
		IsAuth:  true,
	},
}
