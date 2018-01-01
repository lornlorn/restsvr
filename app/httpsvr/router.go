package httpsvr

import (
	"app/models"
	"app/redisctr"
	"app/utils"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"

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
		// time.Sleep(10 * time.Second)
		fmt.Fprintln(res, "Route aaa Finish")
		return
	}
	tmpl, err := template.ParseFiles("views/error/404.html")
	if err != nil {
		log.Printf("Parse Error : %v\n", err)
		return
	}
	tmpl.Execute(res, req.URL)
}

func indexHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Route Index : %v\n", req.URL)
}

func htmlHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route HTML : %v\n", req.URL)
	vars := mux.Vars(req)
	subroute := vars["module"]
	tmpl, err := template.ParseFiles(fmt.Sprintf("views/html/%v.html", subroute))
	if err != nil {
		log.Printf("Parse Error : %v\n", err)
		return
	}
	tmpl.Execute(res, nil)
}

func ajaxHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route Ajax : %v\n", req.URL)
	log.Printf("Request Method : %v\n", req.Method)

	// 获取子路由
	vars := mux.Vars(req)
	subroute := vars["func"]

	// 定义子路由反射调用函数
	ajaxFuncList := map[string]interface{}{
		"addtask": addtask,
	}

	// 获取请求包体(json数据)
	reqBody, _ := ioutil.ReadAll(req.Body)
	// 获取当前用户
	user := gjson.Get(string(reqBody), "user")

	checkresult := utils.CheckAjaxPermission(subroute, user.String())
	if !checkresult {
		fmt.Fprintf(res, "User [%v] permission deny", user.String())
	}

	ret, err := utils.ReflectCall(ajaxFuncList, subroute, user.String())
	if err != nil {
		log.Fatalf("Method [%v] invoke error : %v\n", subroute, err)
	}
	log.Println(ret)

	var retcode, retmsg string

	retcode = "OK"
	retmsg = "成功"

	ajaxres := models.AjaxResMessage{RetCode: retcode, RetMsg: retmsg}
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
