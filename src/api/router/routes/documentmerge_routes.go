package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var documentmergeroutes = []Route{
	Route{
		Uri:     "/documentmerge",
		Method:  http.MethodPost,
		IsAuth:  true,
		Handler: controllers.CreateDocumentMerge,
	},
	Route{
		Uri:     "/documentmerge",
		Method:  http.MethodGet,
		IsAuth:  true,
		Handler: controllers.GetDocumentMerges,
	},
	Route{
		Uri:     "/documentmerge/{id}",
		Method:  http.MethodDelete,
		IsAuth:  true,
		Handler: controllers.DeleteDocumentMerge,
	},
	Route{
		Uri:     "/documentmergelastcode/",
		Method:  http.MethodGet,
		IsAuth:  true,
		Handler: controllers.GetLastDocumentMergecode,
	},
	Route{
		Uri:     "/ledgerdocuments",
		Method:  http.MethodPost,
		IsAuth:  true,
		Handler: controllers.GetLedgerDocumentByid,
	},
}
