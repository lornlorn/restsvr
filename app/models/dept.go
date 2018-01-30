package models

import "time"

type Dept struct {
	DeptId         int       `xorm:"not null pk autoincr unique INTEGER"`
	DeptName       string    `xorm:"not null VARCHAR(16)"`
	DeptStatus     string    `xorm:"not null VARCHAR(16)"`
	DeptCreate     int       `xorm:"not null default 0 INTEGER"`
	DeptCreateTime time.Time `xorm:"not null DATETIME created"`
	DeptModify     int       `xorm:"not null default 0 INTEGER"`
	DeptModifyTime time.Time `xorm:"not null DATETIME updated"`
}
