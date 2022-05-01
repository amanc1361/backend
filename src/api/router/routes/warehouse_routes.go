package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var wareHouseRoute = []Route{
	{
		Uri:     "/warehouse",
		Method:  http.MethodGet,
		Handler: controllers.GetWareHousings,
		IsAuth:  true,
	},
	{
		Uri:     "/warehouse/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetWareHousings,
		IsAuth:  true,
	},
	{
		Uri:     "/warehouse",
		Method:  http.MethodPost,
		Handler: controllers.GetWareHousings,
		IsAuth:  true,
	},

}