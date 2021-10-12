package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var subledgerRoutes = []Route{
	Route{
		Uri:     "/subledger/list",
		Method:  http.MethodPost,
		Handler: controllers.GetSubLedgers,
		IsAuth:  true,
	},
	Route{
		Uri:     "/subledger/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetSubLedger,
		IsAuth:  true,
	},
	Route{
		Uri:     "/subledgersbycode/{code}",
		Method:  http.MethodGet,
		Handler: controllers.GetSubLedgerByCode,
		IsAuth:  false,
	},
	Route{
		Uri:     "/subledgers/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetByLedgerID,
		IsAuth:  false,
	},
	Route{
		Uri:     "/subledgers",
		Method:  http.MethodPost,
		Handler: controllers.CreateSubLedger,
		IsAuth:  true,
	},
	Route{
		Uri:     "/subledgers",
		Method:  http.MethodPut,
		Handler: controllers.UpdateSubLedger,
		IsAuth:  true,
	},
	Route{
		Uri:     "/subledgers/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteSubLedger,
		IsAuth:  true,
	},
	Route{
		Uri:     "/subledger/lastcode/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetLastSubLeadgercode,
		IsAuth:  true,
	},
}
