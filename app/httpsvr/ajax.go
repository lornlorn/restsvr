package httpsvr

import (
	"app/utils"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tidwall/gjson"
)

func addtask(reqBody []byte) []byte {
	var retdata []byte
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
	retdata, _ = utils.Convert2JSON(genResMsg("0000", ""))
	return retdata
}

func test(reqBody []byte) []byte {
	// var resflag string
	resdata := make(map[string]string)
	uid, err := utils.GetUniqueID()
	if err != nil {
		errmsg := fmt.Sprintf("Get Unique ID Failed : %v", err)
		log.Println(errmsg)
		// resflag = "MSG"
		resdata["retcode"] = "0006"
		resdata["retmsg"] = "获取UID失败"
	} else {
		resdata["uid"] = uid
	}
	ret, err := json.Marshal(resdata)
	if err != nil {
		log.Printf("Marshal Json Error : %v\n", err)
	}

	go func(host string) {
		ssh := sshclient.NewSSH("198.211.33.76", "root",
			conf.Password, 22)
		ssh.PrintRun("df -h")

	return ret
}
