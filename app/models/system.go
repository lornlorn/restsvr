package models

import (
	"app/db"
	"fmt"
	"log"
	"time"
)

// System struct map db table system
type System struct {
	SystemId         int       `xorm:"not null pk autoincr unique INTEGER"` // xorm
	SystemCnname     string    `xorm:"not null VARCHAR(128)"`
	SystemEnname     string    `xorm:"not null VARCHAR(16)"`
	SystemType       string    `xorm:"not null VARCHAR(16)"`
	SystemOwner      int       `xorm:"not null default 0 INTEGER"`
	SystemSubowner   int       `xorm:"INTEGER"`
	SystemStatus     string    `xorm:"not null VARCHAR(16)"`
	SystemCreate     int       `xorm:"not null default 0 INTEGER"`
	SystemCreateTime time.Time `xorm:"not null DATETIME created"`
	SystemModify     int       `xorm:"not null default 0 INTEGER"`
	SystemModifyTime time.Time `xorm:"not null DATETIME updated"`
}

// GetSystemList func() ([]System, error)
func GetSystemList(enname string) ([]System, error) {

	systems := make([]System, 0)
	if err := db.Engine.Where("system_status = ? and system_enname like ?", "VALID", enname+"%").Find(&systems); err != nil {
		// return nil, err
		log.Println(err)
	}

	for i, v := range systems {
		fmt.Printf("DataIndex : %v        DataContent : %v\n", i, v)
	}

	ret := []System{
		{
			SystemId:     111,
			SystemEnname: "ORSS",
			SystemCnname: "海外报表平台",
		},
		{
			SystemId:     222,
			SystemEnname: "PCMS",
			SystemCnname: "第三方CA系统",
		},
	}

	return ret, nil

}

// Save insert method
func (s *System) Save() error {
	// affected, err := db.Engine.Insert(d)
	_, err := db.Engine.Insert(s)
	if err != nil {
		return err
	}
	return nil
}
