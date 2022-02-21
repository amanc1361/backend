package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var peopleRoutes = []Route{
	{
		Uri:     "/people/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetPeople,
		IsAuth:  true,
	},
	{
		Uri:     "/people",
		Method:  http.MethodPost,
		Handler: controllers.GetPeoples,
		IsAuth:  true,
	},

	{
		Uri:     "/person",
		Method:  http.MethodPost,
		Handler: controllers.CreatePeople,
		IsAuth:  true,
	},
	{
		Uri:     "/people",
		Method:  http.MethodPut,
		Handler: controllers.UpdatePeople,
		IsAuth:  true,
	},
	{
		Uri:     "/people",
		Method:  http.MethodDelete,
		Handler: controllers.DeletePeople,
		IsAuth:  true,
	},
{
	Uri: "/people/getremperson",
	Method: http.MethodPost,
	Handler: controllers.GetRemPerson,
	IsAuth: true,

},
}
