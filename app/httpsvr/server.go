package httpsvr

import (
	"log"
	"net/http"
)

/*
StartHTTP func()
*/
func StartHTTP() error {
	log.Println("Start HTTP Server...")
	log.Println("-> Initialize HTTP Routes...")
	initRoutes()
	log.Println("-> Listen HTTP Port And Serve...")
	err := http.ListenAndServe(":8888", nil)
	return err
}

func initRoutes() {
	http.HandleFunc("/", notFoundHandler)
	http.HandleFunc("/index", indexHandle)
	http.HandleFunc("/test", testHandle)
	http.HandleFunc("/ajax", ajaxHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))
}
