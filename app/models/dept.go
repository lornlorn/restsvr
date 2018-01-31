package models

import (
	"app/db"
	"time"
)

// Dept struct map db table dept
type Dept struct {
	DeptId         int       `xorm:"not null pk autoincr unique INTEGER"`
	DeptName       string    `xorm:"not null VARCHAR(16)"`
	DeptStatus     string    `xorm:"not null VARCHAR(16)"`
	DeptCreate     int       `xorm:"not null default 0 INTEGER"`
	DeptCreateTime time.Time `xorm:"not null DATETIME created"`
	DeptModify     int       `xorm:"not null default 0 INTEGER"`
	DeptModifyTime time.Time `xorm:"not null DATETIME updated"`
}

// Save insert method
func (d *Dept) Save() error {
	// affected, err := db.Engine.Insert(d)
	_, err := db.Engine.Insert(d)
	if err != nil {
		return err
	}
	return nil
}
