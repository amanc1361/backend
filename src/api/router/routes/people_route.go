package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var peopleRoutes = []Route{
	Route{
		Uri:     "/people/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetPeople,
		IsAuth:  true,
	},
	Route{
		Uri:     "/people",
		Method:  http.MethodPost,
		Handler: controllers.GetPeoples,
		IsAuth:  true,
	},

	Route{
		Uri:     "/person",
		Method:  http.MethodPost,
		Handler: controllers.CreatePeople,
		IsAuth:  true,
	},
	Route{
		Uri:     "/people",
		Method:  http.MethodPut,
		Handler: controllers.UpdatePeople,
		IsAuth:  true,
	},
	Route{
		Uri:     "/people",
		Method:  http.MethodDelete,
		Handler: controllers.DeletePeople,
		IsAuth:  true,
	},
}
