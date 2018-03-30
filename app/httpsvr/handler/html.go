package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HTMLHandler func(res http.ResponseWriter, req *http.Request)
func HTMLHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route HTML : %v\n", req.URL)
	vars := mux.Vars(req)
	group := vars["group"]
	module := vars["module"]
	tmpl, err := template.ParseFiles(fmt.Sprintf("views/html/%v/%v.html", group, module))
	if err != nil {
		log.Printf("Parse Error : %v\n", err)

		//二级子路由模板不存在返回404页面
		tmpl, err := template.ParseFiles("views/error/404.html")
		if err != nil {
			log.Printf("Parse Error : %v\n", err)
			return
		}
		tmpl.Execute(res, req.URL)

		return
	}

	tmpl.Execute(res, nil)
}
