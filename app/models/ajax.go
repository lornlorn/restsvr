package models

/*
AjaxResMessage struct
*/
type AjaxResMessage struct {
	RetCode string      `json:"retcode"`
	RetMsg  string      `json:"retmsg"`
	RetData interface{} `json:"data"`
}

/*
AjaxCompleteSystem struct
*/
type AjaxCompleteSystem struct {
	SysID     int    `json:"sysid"`
	SysEnName string `json:"sysenname"`
	SysCnName string `json:"syscnname"`
}
