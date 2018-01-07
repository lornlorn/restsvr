package httpsvr

import (
	"log"
)

func addtask(reqBody []byte) bool {
	log.Println("ajax.addtask")
	log.Println(string(reqBody))
	return true
}
