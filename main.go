package main

import (
	"app/db"
	"app/httpsvr"
	"log"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(httpsvr.StartHTTP())
	_ = db.Engine.Close()
}
