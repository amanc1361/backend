package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var storepersonreciveRoutes = []Route{
	Route{
		Uri:     "/storepersonrecive/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetStorePersonRecive,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storepersonrecives",
		Method:  http.MethodPost,
		Handler: controllers.GetStorePersonRecives,
		IsAuth:  true,
	},

	Route{
		Uri:     "/storepersonrecive",
		Method:  http.MethodPost,
		Handler: controllers.CreateStorePersonRecive,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storepersonrecive/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateStorePersonRecive,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storepersonrecive/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteStorePersonRecive,
		IsAuth:  true,
	},
}
