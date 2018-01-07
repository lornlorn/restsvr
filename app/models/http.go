package models

/*
AjaxReq struct
*/
type AjaxReq struct {
	Module string      `json:"module"`
	User   string      `json:"user"`
	Data   interface{} `json:"data"`
}

/*
AjaxResMessage struct
*/
type AjaxResMessage struct {
	RetCode string `json:"retcode"`
	RetMsg  string `json:"retmsg"`
}

/*
AjaxResData struct
*/
type AjaxResData struct {
	RetData interface{} `json:"retdata"`
}
