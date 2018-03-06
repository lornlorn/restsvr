package api

import (
	"app/models"
	"app/utils"
	"log"
)

// GetSystemList funct(reqBody []byte) ([]byte, error)
func GetSystemList(reqBody []byte) ([]byte, error) {
	syslist, err := models.GetSystemList()
	if err != nil {
		log.Printf("Get System List Fail : %v\n", err)
		return nil, err
	}

	log.Println(len(syslist))
	var ACS [len(syslist)]models.AjaxCompleteSystem
	for i, v := range syslist {
		log.Printf("index : %v, %v", i, v)
		ACS[i] = models.AjaxCompleteSystem{
			SysID:     v.SystemId,
			SysEnName: v.SystemEnname,
			SysCnName: v.SystemCnname,
		}
	}

	ret, err := utils.Convert2JSON(ACS)
	if err != nil {
		log.Printf("Marshal Json Error : %v\n", err)
		return nil, err
	}
	return ret, nil
}
