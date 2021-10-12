package routes

import (
	"back-account/src/api/controllers"
	"net/http"
)

var documentRoutes = []Route{

	Route{
		Uri:     "/document",
		Method:  http.MethodPost,
		Handler: controllers.CreateDocument,
		IsAuth:  true,
	},
	Route{
		Uri:     "/document/list",
		Method:  http.MethodPost,
		Handler: controllers.GetDocuments,
		IsAuth:  true,
	},
	Route{
		Uri:     "/document/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetDocument,
		IsAuth:  true,
	},
	Route{
		Uri:     "/documentbycode/{code}",
		Method:  http.MethodPost,
		Handler: controllers.GetDocumentByCode,
		IsAuth:  true,
	},
	Route{
		Uri:     "/document/des/des",
		Method:  http.MethodGet,
		Handler: controllers.GetDocumentByDescription,
		IsAuth:  true,
	},
	Route{
		Uri:     "/document/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateDocument,
		IsAuth:  true,
	}, Route{
		Uri:     "/documentlastcode",
		Method:  http.MethodPost,
		Handler: controllers.GetLastDocumentcode,
		IsAuth:  true,
	},
	Route{
		Uri:     "/sortdocument",
		Method:  http.MethodPost,
		Handler: controllers.SortDocuments,
		IsAuth:  true,
	},
	Route{
		Uri:     "/documentrows/{id}",
		Method:  http.MethodPost,
		Handler: controllers.GetDocumentRowsbyDocumentid,
		IsAuth:  true,
	},
	Route{
		Uri:     "/document/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteDocument,
		IsAuth:  true,
	},
	Route{
		Uri:     "/firstdate",
		Method:  http.MethodPost,
		Handler: controllers.GetFirstDate,
		IsAuth:  true,
	},
	Route{
		Uri:     "/lastdate",
		Method:  http.MethodPost,
		Handler: controllers.GetLasttDate,
		IsAuth:  true,
	},
	Route{
		Uri:     "/documentinfo",
		Method:  http.MethodPost,
		Handler: controllers.GetDocInfo,
		IsAuth:  true,
	},
	Route{
		Uri:     "/documentbytype",
		Method:  http.MethodPost,
		Handler: controllers.GetDcoumentsByTypeDoc,
		IsAuth:  true,
	},
}
