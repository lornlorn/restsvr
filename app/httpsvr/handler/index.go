package handler

import (
	"html/template"
	"log"
	"net/http"
)

// IndexHandle func(res http.ResponseWriter, req *http.Request)
func IndexHandle(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route Index : %v\n", req.URL)
	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Printf("Parse Error : %v\n", err)
		return
	}

	tmpl.Execute(res, nil)
}
