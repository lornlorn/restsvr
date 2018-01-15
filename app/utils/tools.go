package utils

import (
	"encoding/json"
	"log"
)

// Convert2JSON (data interface{}) []byte
func Convert2JSON(data interface{}) ([]byte, error) {
	var retdata []byte
	var err error
	switch data.(type) {
	case []byte:
		go func(b []byte) {
			log.Println("Response JSON Data :")
			str := string(b[:])
			log.Println(str)
		}(data.([]byte))
		// log.Printf("Response Data : %v", data)
		retdata = data.([]byte)
	default:
		log.Printf("Response Data To JSON : %v", data)
		retdata, err = json.Marshal(data)
		if err != nil {
			log.Printf("Marshal Json Error : %v\n", err)
		}
	}
	return retdata, err
}
