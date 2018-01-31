package db

import (
	"database/sql"
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

// MyDB struct
type MyDB struct {
	DBType string
	*sql.DB
}

const (
	host     = "192.168.100.100"
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
	dbstr := "postgres://test:test@192.168.100.100:5432/testdb?sslmode=disable"
	Engine, err = xorm.NewEngine(dbtype, dbstr)
	if err != nil {
		return err
	}
	Engine.ShowSQL(true)
	err = Engine.Ping()
	if err != nil {
		log.Printf("DB Ping Failed : %v\n", err)
		return err
	}
	return nil
}
