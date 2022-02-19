package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var documentmergeroutes = []Route{
	{
		Uri:     "/documentmerge",
		Method:  http.MethodPost,
		IsAuth:  true,
		Handler: controllers.CreateDocumentMerge,
	},
	{
		Uri:     "/documentmerge",
		Method:  http.MethodGet,
		IsAuth:  true,
		Handler: controllers.GetDocumentMerges,
	},
	{
		Uri:     "/documentmerge/{id}",
		Method:  http.MethodDelete,
		IsAuth:  true,
		Handler: controllers.DeleteDocumentMerge,
	},
	{
		Uri:     "/documentmergelastcode/",
		Method:  http.MethodGet,
		IsAuth:  true,
		Handler: controllers.GetLastDocumentMergecode,
	},
	{
		Uri:     "/ledgerdocuments",
		Method:  http.MethodPost,
		IsAuth:  true,
		Handler: controllers.GetLedgerDocumentByid,
	},
}
