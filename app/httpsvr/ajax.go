package httpsvr

import (
	"log"
)

func addtask(reqBody []byte) (string, string) {
	log.Println("ajax.addtask")
	log.Println(reqBody)
	return "ok", ""
}
