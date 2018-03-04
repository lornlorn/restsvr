package handler

import (
	"app/models"
	"app/utils"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AutocompleteHandler func(res http.ResponseWriter, req *http.Request)
func AutocompleteHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Router URL : %v\n", req.URL)
	vars := mux.Vars(req)
	subroute := vars["func"]

	// 获取请求包体(json数据)
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Request Body Read Failed : %v\n", err)
		ResponseAjaxMsg(res, "2000", nil)
		return
	}
	log.Println("Request JSON Content :")
	log.Println(string(reqBody))

	switch subroute {
	case "systemlist":
		models.GetSystemList
	default:

	}

	tmplData := reqTmplData(subroute)
	ret, err := utils.Convert2JSON(tmplData)
	if err != nil {
		log.Printf("Autocomplete Error\n")
		ResponseAjaxMsg(res, "2000", nil)
		return
	}
	res.Write(ret)

}

// DataHandler func(res http.ResponseWriter, req *http.Request)
func DataHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route HTML : %v\n", req.URL)
	vars := mux.Vars(req)
	subroute := vars["module"]

}
