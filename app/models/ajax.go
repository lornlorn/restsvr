package models

/*
AjaxResMessage struct
*/
type AjaxResMessage struct {
	RetCode string      `json:"retcode"`
	RetMsg  string      `json:"retmsg"`
	RetData interface{} `json:"data"`
}
