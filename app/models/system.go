package models

import (
	"time"
)

// System struct map db table system
type System struct {
	SystemId         int       `xorm:"not null pk autoincr unique INTEGER"`
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

// GetSystemList func() []System
func GetSystemList() []System {
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

	return ret
}
