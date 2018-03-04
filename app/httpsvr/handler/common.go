package handler

import (
	"app/models"
	"app/utils"
	"log"
	"net/http"
)

// test
func reqTmplData(module string) map[string]string {
	ret := map[string]string{
		"111": "PCMS-第三方CA系统",
		"222": "GLMS-全球额度管理系统",
		"333": "GLS-总账系统",
		"444": "ORSS-海外报表平台",
	}
	return ret
}

// ResponseAjaxMsg func(resWriter http.ResponseWriter, retcode string, retmsg string)
func ResponseAjaxMsg(resWriter http.ResponseWriter, retcode string, retdata interface{}) {
	retmsg := utils.GetRetMsg(retcode)

	arm := models.AjaxResMessage{
		RetCode: retcode,
		RetMsg:  retmsg,
		RetData: retdata,
	}
	ret, err := utils.Convert2JSON(arm)
	if err != nil {
		log.Printf("Marshal Json Error : %v\n", err)
	}
	resWriter.Write(ret)
}
