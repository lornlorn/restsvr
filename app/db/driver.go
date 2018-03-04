package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq" // postgresql driver
)

// MyDB struct
type MyDB struct {
	DBType string
	*sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "testdb"
)

// Engine xorm
var Engine *xorm.Engine

// InitDB func() error
func InitDB() error {
	var err error
	dbtype := "postgres"
	// dbstr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", host, port, user, password, dbname)
	// dbstr := "postgres://test:test@localhost:5432/testdb?sslmode=disable"
	dbstr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)
	Engine, err = xorm.NewEngine(dbtype, dbstr)
	if err != nil {
		return err
	}
	Engine.ShowSQL(true)
	err = Engine.Ping()
	if err != nil {
		log.Println("DB Ping Failed")
		return err
	}
	return nil
}
