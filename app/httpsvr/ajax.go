package httpsvr

import (
	"encoding/json"
	"log"

	"github.com/tidwall/gjson"
)

func addtask(reqBody []byte) (string, string, string) {
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
	return "MSG", "0000", ""
}

func test(reqBody []byte) (string, []byte) {
	log.Println("请求JSON正文:")
	log.Println(string(reqBody))
	res := map[string]string{"aaa": "test"}
	ret, err := json.Marshal(res)
	if err != nil {
		log.Printf("Marshal Json Error : %v\n", err)
	}
	return "JSON", ret
}
