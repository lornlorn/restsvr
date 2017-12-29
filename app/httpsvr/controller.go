package httpsvr

import (
	"app/models"
	"app/redisctr"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
			// time.Sleep(10 * time.Second)
		}
		time.Sleep(10 * time.Second)
		fmt.Fprintln(res, "Route aaa Finish")
		return
	}
	tmpl, err := template.ParseFiles("views/error/404.html")
	if err != nil {
		log.Printf("Parse Error : %v\n", err)
	}
	tmpl.Execute(res, req.URL)
}

func indexHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Route Index : %v\n", req.URL)
}

func htmlHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route HTML : %v\n", req.URL)
	vars := mux.Vars(req)
	tmpl, err := template.ParseFiles(fmt.Sprintf("views/html/%v.html", vars["module"]))
	if err != nil {
		log.Printf("Parse Error : %v\n", err)
	}
	tmpl.Execute(res, nil)
}

func ajaxHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route Ajax : %v\n", req.URL)
	log.Printf("Request Method : %v\n", req.Method)
	vars := mux.Vars(req)
	subroute := vars["func"]

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
}

func testHandle(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route Test : %v\n", req.URL)
	tmpl, err := template.ParseFiles("views/test/test.html")
	if err != nil {
		log.Printf("Parse Error : %v\n", err)
	}
	tmpl.Execute(res, nil)
}

func redisHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route Redis : %v\n", req.URL)
	connStr := models.RedisConnector{
		Proto: "tcp",
		Addr:  "127.0.0.1",
		Port:  6379,
	}
	redisConn, err := redisctr.Connect(connStr)
	if err != nil {
		log.Printf("Redis Connect Failed : %v\n", err)
	}

	redisClient := redisctr.RedisClient{Client: redisConn}
	// err = redisController.Set("username", "hqdiaolei")
	// if err != nil {
	// 	log.Printf("Redis SET Failed : %v\n", err)
	// }

	// reply, err := redisController.Get("username")
	// if err != nil {
	// 	log.Printf("Redis GET Failed : %v\n", err)
	// }
	err = redisClient.Lpush("chan", "task1")
	if err != nil {
		log.Printf("Redis LPUSH Failed : %v\n", err)
	}
	err = redisClient.Lpush("chan", "task2")
	if err != nil {
		log.Printf("Redis LPUSH Failed : %v\n", err)
	}

	reply, err := redisClient.Rpop("chan")
	if err != nil {
		log.Printf("Redis RPOP Failed : %v\n", err)
	}

	err = redisClient.Close()
	if err != nil {
		log.Printf("Redis Connector Close Failed : %v\n", err)
	}
	fmt.Fprintf(res, "Route Redis Finish : %v", reply)
}
