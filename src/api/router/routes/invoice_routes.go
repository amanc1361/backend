package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var invoiceRoutes = []Route{
	{
		Uri:     "/invoic",
		Method:  http.MethodPost,
		Handler: controllers.CreateInvocie,
		IsAuth:  true,
	},
	{
		Uri:     "/invoic",
		Method:  http.MethodGet,
		Handler: controllers.GetInovices,
		IsAuth:  true,
	},
	{
		Uri:     "/invoic/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetInvoice,
		IsAuth:  true,
	},
}