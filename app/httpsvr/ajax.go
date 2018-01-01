package httpsvr

import (
	"log"
)

func addtask(req string) string {
	log.Println("ajax.addtask")
	log.Println(req)
	return "ok"
}
