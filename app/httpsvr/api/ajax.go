package api

import (
	"app/models"
	"app/utils"
	"log"

	"github.com/tidwall/gjson"
)

// GetSystemList funct(reqBody []byte) ([]byte, error)
func GetSystemList(reqBody []byte) ([]byte, error) {
	keyword := gjson.Get(string(reqBody), "data.keyword")
	syslist, err := models.GetSystemList(keyword.String())
	if err != nil {
		log.Printf("Get System List Fail : %v\n", err)
		return nil, err
	}

	// ajaxCompleteSystem := make([]models.AjaxCompleteSystem, 0, len(syslist))
	ajaxCompleteSystem := make([]models.AjaxCompleteSystem, len(syslist))
	for i, v := range syslist {
		// log.Printf("index : %v, %v", i, v)
		ajaxCompleteSystem[i] = models.AjaxCompleteSystem{
			SysID:     v.SystemId,
			SysEnName: v.SystemEnname,
			SysCnName: v.SystemCnname,
		}
	}

	ret, err := utils.Convert2JSON(ajaxCompleteSystem)
	if err != nil {
		log.Printf("Marshal Json Error : %v\n", err)
		return nil, err
	}
	return ret, nil
}
