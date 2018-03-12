package models

import (
	"app/db"
	"log"
	"strings"
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
func GetSystemList(enkeyword string) ([]System, error) {

	systems := make([]System, 0)
	if err := db.Engine.Where("system_status = ? and upper(system_enname) like ?", "VALID", strings.ToUpper(enkeyword)+"%").Find(&systems); err != nil {
		// return nil, err
		log.Println(err)
		return nil, err
	}

	for i, v := range systems {
		log.Printf("DataIndex : %v, DataContent : %v\n", i, v)
	}

	return systems, nil
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
