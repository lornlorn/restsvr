package utils

import (
	"encoding/json"
	"log"
)

// Convert2JSON (data interface{}) []byte
func Convert2JSON(data interface{}) ([]byte, error) {
	switch data.(type) {
	case []byte:
		// log.Println("Convert To JSON args []byte")
		retdata := data.([]byte)
		return retdata, nil
	default:
		// log.Println("Convert To JSON args not []byte")
		retdata, err := json.Marshal(data)
		if err != nil {
			log.Printf("Marshal Json Error : %v\n", err)
		}
		return retdata, err
	}
}
