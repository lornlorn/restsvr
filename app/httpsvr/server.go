package httpsvr

import (
	"app/httpsvr/handler"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/*
StartHTTP func()
*/
func StartHTTP() error {
	log.Println("Start HTTP Server...")
	log.Println("-> Initialize HTTP Routes...")
	r := mux.NewRouter()
	initRoutes(r)
	log.Println("-> Listen HTTP Port And Serve...")
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8888",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	// err := http.ListenAndServe(":8888", r)
	err := srv.ListenAndServe()
	return err
}

func initRoutes(r *mux.Router) {
	// normal router
	r.HandleFunc("/index", handler.IndexHandle)

	// dynamic router
	// ajax router
	ajax := r.PathPrefix("/ajax").Subrouter()
	// ajax root router
	// ajax.HandleFunc("/", notFoundHandler)
	// ajax autocomplete subrouter
	ajax.HandleFunc("/autocomplete/{func}", handler.AutocompleteHandler)
	// ajax data subrouter
	ajax.HandleFunc("/data/{func}", handler.DataHandler)

	// html router
	h := r.PathPrefix("/html").Subrouter()
	h.HandleFunc("/", handler.NotFoundHandler)
	h.HandleFunc("/{key}", handler.NotFoundHandler)
	h.HandleFunc("/{group}/{module}", handler.HTMLHandler)

	// test
	r.HandleFunc("/redis", handler.RedisHandler)
	r.HandleFunc("/test/{page}", handler.TestHandle)

	// static resource
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// subrouter
	root := r.PathPrefix("/").Subrouter()
	root.HandleFunc("/", handler.NotFoundHandler)
	root.HandleFunc("/{key}", handler.NotFoundHandler)

	// http.HandleFunc("/", notFoundHandler)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
