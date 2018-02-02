package main

import (
	"app/db"
	"app/httpsvr"
	"log"
)

func main() {
	log.Println("Init DB Connect...")
	err := db.InitDB()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB Connect Success...")
	log.Fatalln(httpsvr.StartHTTP())
	_ = db.Engine.Close()
}
