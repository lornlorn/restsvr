package httpsvr

import (
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
	r.HandleFunc("/index", indexHandle)

	r.HandleFunc("/ajax/{func}", ajaxHandler)
	r.HandleFunc("/html/{module}", htmlHandler)

	r.HandleFunc("/redis", redisHandler)
	r.HandleFunc("/test", testHandle)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	s := r.PathPrefix("/").Subrouter()
	s.HandleFunc("/", notFoundHandler)
	s.HandleFunc("/{key}", notFoundHandler)

	// http.HandleFunc("/", notFoundHandler)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
