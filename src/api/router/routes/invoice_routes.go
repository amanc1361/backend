package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var invoiceRoutes = []Route{
	Route{
		Uri:     "/invoice",
		Method:  http.MethodPost,
		Handler: controllers.CreateInvocie,
		IsAuth:  true,
	},
	Route{
		Uri:     "/invoice",
		Method:  http.MethodPut,
		Handler: controllers.UpdateInvoice,
		IsAuth:  true,
	},
	Route{
		Uri:     "/invoice",
		Method:  http.MethodGet,
		Handler: controllers.GetInovices,
		IsAuth:  true,
	},
	Route{
		Uri:     "/invoice/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetInvoice,
		IsAuth:  true,
	},
	Route{
		Uri:     "/invoice/selltype",
		Method:  http.MethodPost,
		Handler: controllers.GetSellTypies,
		IsAuth:  true,
	},
	Route{
		Uri:     "/invoice/getlastnumber",
		Method:  http.MethodPost,
		Handler: controllers.GetInvoiceNumber,
		IsAuth:  true,
	},
	Route{
		Uri:     "/invoice/taxreport",
		Method:  http.MethodGet,
		Handler: controllers.GetTaxYear,
		IsAuth:  true,
	},
	
}