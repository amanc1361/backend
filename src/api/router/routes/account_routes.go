package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var accountRoutes = []Route{
	{
		Uri:    "/account",
		Method: http.MethodPost,
		Handler: controllers.SaveAccount,
		IsAuth: true,
	},
	{
		Uri:    "/account",
		Method: http.MethodPut,
		Handler: controllers.SaveAccount,
		IsAuth: true,
	},
	{
		Uri:    "/account/{accountid}",
		Method: http.MethodGet,
		Handler: controllers.GetAccount,
		IsAuth: true,
	},
	{
		Uri:    "/account",
		Method: http.MethodGet,
		Handler: controllers.GetAccounts,
		IsAuth: true,
	},
}