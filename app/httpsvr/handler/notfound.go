package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*
Route Not Found 404 Page
And
Route "/" Direct To "/index"
*/
func NotFoundHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		http.Redirect(res, req, "/index", http.StatusFound)
	} else if req.URL.Path == "/aaa" {
		for i := 0; i <= 10; i++ {
			go fmt.Printf("%v %v\n", req.URL, i)
			// time.Sleep(10 * time.Second)
		}
		// time.Sleep(10 * time.Second)
		fmt.Fprintln(res, "Route aaa Finish")
		return
	}
	tmpl, err := template.ParseFiles("views/error/404.html")
	if err != nil {
		log.Printf("Parse Error : %v\n", err)
		return
	}
	tmpl.Execute(res, req.URL)
}
