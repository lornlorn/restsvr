package httpsvr

import (
	"app/models"
	"app/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"

	"github.com/gorilla/mux"
)

func ajaxHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route Ajax : %v\n", req.URL)
	log.Printf("Request Method : %v\n", req.Method)

	// 获取子路由
	vars := mux.Vars(req)
	subroute := vars["module"]

	// 定义子路由反射调用函数
	ajaxFuncList := map[string]interface{}{
		"taskadd": taskadd,
		"test":    test,
		// "autocomplete": autocomplete,
	}

	// subroute()

	// 定义返回值
	// var retcode, retmsg string

	// 获取请求包体(json数据)
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errmsg := fmt.Sprintf("Request Body Read Failed : %v\n", err)
		log.Println(errmsg)
		retdata := genResMsg("0001", errmsg)
		ajaxResponse(res, retdata)
		return
	}
	log.Println("Request JSON Content :")
	log.Println(string(reqBody))
	// 获取当前用户
	user := gjson.Get(string(reqBody), "user")

	// 调用Ajax权限检查
	checkresult := utils.CheckAjaxPermission(subroute, user.String())
	if !checkresult {
		errmsg := fmt.Sprintf("User [%v] Permission Deny", user.String())
		log.Println(errmsg)
		retdata := genResMsg("0002", errmsg)
		ajaxResponse(res, retdata)
		return
	}

	// 调用反射方法
	ret, err := utils.ReflectCall(ajaxFuncList, subroute, reqBody)
	if err != nil {
		errmsg := fmt.Sprintf("Method [%v] Invoke Error : %v\n", subroute, err)
		log.Println(errmsg)
		retdata := genResMsg("0003", errmsg)
		ajaxResponse(res, retdata)
		return
	}
	// 处理反射方法返回值
	// 规定必须两个返回值
	// 1、string类型("data","msg")表示返回数据为数据或是消息
	// 2、JSON格式数据([]byte())
	switch len(ret) {
	case 2:
		ajaxResponse(res, ret[1].Bytes())
	default:
		retdata := genResMsg("0005", "反射方法返回值数量不正确")
		ajaxResponse(res, retdata)
	}

	return

	// 成功应答
	// retdata := genResMsg("0000", "成功")
	// ajaxResponse(res, retdata)
}

// Ajax Message {RetCode, RetMsg} To JSON
func genResMsg(retcode string, retmsg string) models.AjaxResMessage {
	ajaxres := models.AjaxResMessage{RetCode: retcode, RetMsg: retmsg}
	return ajaxres
}

// Ajax Return Data To JSON
/* Unused
func ajaxDataToJSON(data interface{}) []byte {
	var retdata []byte
	var err error
	switch data.(type) {
	case []byte:
		retdata = data.([]byte)
	default:
		retdata, err = json.Marshal(data)
		if err != nil {
			log.Printf("Marshal Json Error : %v\n", err)
		}
	}
	return retdata
}
*/

// AjaxResponse
func ajaxResponse(response http.ResponseWriter, retobj interface{}) {
	var retdata []byte
	var err error
	switch retobj.(type) {
	case []byte:
		go func(b []byte) {
			log.Println("Response JSON Data :")
			str := string(b[:])
			log.Println(str)
		}(retobj.([]byte))
		// log.Printf("Response Data : %v", retobj)
		retdata = retobj.([]byte)
	default:
		log.Printf("Response Data To JSON : %v", retobj)
		retdata, err = json.Marshal(retobj)
		if err != nil {
			log.Printf("Marshal Json Error : %v\n", err)
		}
	}
	response.Write(retdata)
}
