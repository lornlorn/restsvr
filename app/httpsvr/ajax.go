package httpsvr

import (
	"app/models"
	"app/utils"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tidwall/gjson"
)

func taskadd(reqBody []byte) (string, []byte) {
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
	return "msg", retdata
}

func test(reqBody []byte) (string, []byte) {
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

	// go func(host string) {
	// 	ssh := sshclient.NewSSH(host, "root", "", 22)
	// 	ssh.PrintRun("df -h")
	// }("198.211.33.76")

	dept := new(models.Dept)
	dept.DeptName = "TestDept"
	dept.DeptStatus = "VALID"

	err = dept.Save()
	if err != nil {
		log.Printf("DB Save Failed : %v\n", err)
	}
	log.Println(dept.DeptId)

	return "msg", ret
}

func autocomplete(reqBody []byte) (string, []byte) {
	resdata := []map[string]string{
		{"id": "111",
			"enname": "PCMS",
			"cnname": "第三方CA系统",
		},
		{"id": "222",
			"enname": "ORSS",
			"cnname": "海外报表平台",
		},
	}
	ret, err := json.Marshal(resdata)
	if err != nil {
		log.Printf("Marshal Json Error : %v\n", err)
	}
	return "", ret
}
