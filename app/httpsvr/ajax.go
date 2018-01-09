package httpsvr

import (
	"log"

	"github.com/tidwall/gjson"
)

func addtask(reqBody []byte) bool {
	log.Println("请求JSON正文:")
	log.Println(string(reqBody))
	system := gjson.Get(string(reqBody), "data.jobinfo.system")
	steplist := gjson.Get(string(reqBody), "data.steplist")
	steplist.ForEach(func(key, value gjson.Result) bool {
		log.Printf("step : %v\n", key)
		stepdtlList := value
		stepdtlList.ForEach(func(key, value gjson.Result) bool {
			log.Printf("stepdtl : %v\n", key)
			log.Printf("stepdtl info : %v\n", value.String())
			return true
		})
		// log.Println(value.String())
		return true
	})
	log.Printf("系统名称 : %v\n", system.String())
	return true
}
