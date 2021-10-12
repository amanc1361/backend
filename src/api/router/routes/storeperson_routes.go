package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var storepersonRoutes = []Route{
	Route{
		Uri:     "/storeperson/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetStorePerson,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storepersons",
		Method:  http.MethodPost,
		Handler: controllers.GetStorePersons,
		IsAuth:  true,
	},

	Route{
		Uri:     "/storeperson",
		Method:  http.MethodPost,
		Handler: controllers.CreateStorePerson,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeperson/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateStorePerson,
		IsAuth:  true,
	},
	Route{
		Uri:     "/storeperson/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteStorePerson,
		IsAuth:  true,
	},
}
