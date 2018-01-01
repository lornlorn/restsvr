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
AjaxRes struct
*/
type AjaxResMessage struct {
	RetCode string `json:"retcode"`
	RetMsg  string `json:"retmsg"`
}
