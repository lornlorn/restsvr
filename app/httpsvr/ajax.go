package httpsvr

// func ajaxHandler(res http.ResponseWriter, req *http.Request) {
// 	log.Printf("Route Ajax : %v\n", req.URL)
// 	log.Printf("Request Method : %v\n", req.Method)
// 	vars := mux.Vars(req)
// 	subroute := vars["func"]

// 	var ajaxreq models.AjaxReq
// 	reqBody, _ := ioutil.ReadAll(req.Body)

// 	err := json.Unmarshal(reqBody, &ajaxreq)
// 	defer req.Body.Close()
// 	if err != nil {
// 		log.Printf("Unmarshal Json Error : %v\n", err)
// 		return
// 	}

// 	// Show Request JSON Data
// 	log.Println(string(reqBody))
// 	log.Printf("Request Data To Struct : %v", ajaxreq)

// 	var retcode, retmsg string

// 	if ajaxreq.Data["username"] == "hqdiaolei" && ajaxreq.Data["password"] == "123456Aa" {
// 		retcode = "OK"
// 		retmsg = "成功"
// 	} else {
// 		retcode = "FAIL"
// 		retmsg = "失败"
// 	}

// 	ajaxres := models.AjaxRes{RetCode: retcode, RetMsg: retmsg}
// 	log.Printf("Response Data To Struct : %v", ajaxres)

// 	retdata, err := json.Marshal(ajaxres)
// 	if err != nil {
// 		log.Printf("Marshal Json Error : %v\n", err)
// 	}

// 	res.Write(retdata)
// }
