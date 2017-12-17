package models

/*
AjaxReq struct
*/
type AjaxReq struct {
	Module string            `json:"module"`
	Data   map[string]string `json:"data"`
}

/*
AjaxRes struct
*/
type AjaxRes struct {
	RetCode string `json:"retcode"`
	RetMsg  string `json:"retmsg"`
}
