package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TestHandle func(res http.ResponseWriter, req *http.Request)
func TestHandle(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route Test : %v\n", req.URL)
	vars := mux.Vars(req)
	subroute := vars["page"]
	tmpl, err := template.ParseFiles(fmt.Sprintf("views/test/%v.html", subroute))
	if err != nil {
		log.Printf("Parse Error : %v\n", err)
		return
	}
	tmpl.Execute(res, nil)

}
