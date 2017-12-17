package httpsvr

import (
	"app/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

/*
Route Not Found 404 Page
And
Route "/" Direct To "/index"
*/
func notFoundHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		http.Redirect(res, req, "/index", http.StatusFound)
	} else if req.URL.Path == "/aaa" {
		for i := 0; i <= 10; i++ {
			go fmt.Printf("%v %v\n", req.URL, i)
			// time.Sleep(1 * time.Second)
		}
		return
	}
	tmpl, err := template.ParseFiles("views/error/404.html")
	if err != nil {
		log.Printf("Parse Error : %v\n", err)
	}
	tmpl.Execute(res, req.URL)
}

func indexHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "route index : %v\n", req.URL)
}

func testHandle(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route Test : %v\n", req.URL)
	tmpl, err := template.ParseFiles("views/test/test.html")
	if err != nil {
		log.Printf("Parse Error : %v\n", err)
	}
	tmpl.Execute(res, nil)
}

func ajaxHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route Ajax : %v\n", req.URL)
	log.Printf("Request Method : %v\n", req.Method)

	var ajaxreq models.AjaxReq
	reqBody, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(reqBody, &ajaxreq)
	defer req.Body.Close()
	if err != nil {
		log.Printf("Unmarshal Json Error : %v\n", err)
		return
	}

	// Show Request JSON Data
	log.Println(string(reqBody))
	log.Printf("Request Data To Struct : %v", ajaxreq)

	// for k, v := range ajaxreq.Data {
	// 	log.Printf("key = %v, value = %v", k, v)
	// }
	// log.Println(ajaxreq.Data["username"])
	// log.Println(ajaxreq.Data["password"])

	var retcode, retmsg string

	if ajaxreq.Data["username"] == "hqdiaolei" && ajaxreq.Data["password"] == "123456Aa" {
		retcode = "OK"
		retmsg = "成功"
	} else {
		retcode = "FAIL"
		retmsg = "失败"
	}

	ajaxres := models.AjaxRes{RetCode: retcode, RetMsg: retmsg}
	log.Printf("Response Data To Struct : %v", ajaxres)

	retdata, err := json.Marshal(ajaxres)
	if err != nil {
		log.Printf("Marshal Json Error : %v\n", err)
	}
	res.Write(retdata)
	// json.NewEncoder(res).Encode(ret)
}
