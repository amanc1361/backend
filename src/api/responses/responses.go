package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	mydata := result{
		Status: statusCode,
		Mydata: data,
	}
	// console.Preety(mydata)
	err := json.NewEncoder(w).Encode(mydata)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}

}
func JSONWITHPAGENUMBER(w http.ResponseWriter, statusCode int, data interface{}, pagesize uint) {
	w.WriteHeader(statusCode)
	mydata := result{
		Status:   statusCode,
		Mydata:   data,
		Pagesize: pagesize,
	}
	err := json.NewEncoder(w).Encode(mydata)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}

}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {

		JSON(w, statusCode, struct {
			ERROR string `json:"error"`
		}{ERROR: err.Error()})

		return
	}

	JSON(w, http.StatusBadRequest, nil)

}

type result struct {
	Status   int         `json:"status"`
	Mydata   interface{} `json:"data"`
	Pagesize uint        `json:pagesize`
}
