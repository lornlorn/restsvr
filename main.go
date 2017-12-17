package main

import (
	"app/httpsvr"
	"log"
)

func main() {
	log.Fatalln(httpsvr.StartHTTP())
}
