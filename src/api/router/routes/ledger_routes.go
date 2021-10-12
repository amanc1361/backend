package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var ledgerRoutes = []Route{
	Route{
		Uri:     "/ledger/list",
		Method:  http.MethodPost,
		Handler: controllers.GetLedgers,
		IsAuth:  true,
	},
	Route{
		Uri:     "/ledger/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetLedger,
		IsAuth:  true,
	},
	Route{
		Uri:     "/ledger/code/{code}",
		Method:  http.MethodGet,
		Handler: controllers.GetLedgerByCode,
		IsAuth:  true,
	},
	Route{
		Uri:     "/ledgers/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetByGroupID,
		IsAuth:  false,
	},
	Route{
		Uri:     "/ledgers",
		Method:  http.MethodPost,
		Handler: controllers.CreateLedger,
		IsAuth:  true,
	},
	Route{
		Uri:     "/ledgers",
		Method:  http.MethodPut,
		Handler: controllers.UpdateLedger,
		IsAuth:  true,
	},
	Route{
		Uri:     "/ledgers/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteLedger,
		IsAuth:  true,
	},
	Route{
		Uri:     "/ledger/lastcode/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetLastLeadgercode,
		IsAuth:  true,
	},
}
