package handler

import (
	"app/models"
	"app/utils"
	"log"
	"net/http"
)

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
