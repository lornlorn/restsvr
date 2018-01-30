package db

import (
	"database/sql"
	"fmt"
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

var Engine *xorm.Engine

// Connect func()
func Connect() (*MyDB, error) {
	dbstr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbstr)
	if err != nil {
		log.Printf("DB Connect Failed : %v", err)
		return nil, err
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("DB Ping Failed : %v", err)
		return nil, err
	}

	mydb := &MyDB{DB: db}
	return mydb, nil
}

// DBClose (){}
func (db *MyDB) DBClose() {
	db.Close()
}

// Insert (tab string, args []interface{}) error
func (db *MyDB) Insert(tab string, stat map[string]interface{}) error {
	stmt, err := db.Prepare("INSERT INTO test.TB_TEST(id,name) VALUES($1,$2) RETURNING id")
	if err != nil {
		log.Printf("SQL Prepare Failed : %v", err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(1, "test1")
	if err != nil {
		log.Printf("SQL Execute Failed : %v", err)
	}
	log.Println(res)
	return nil
}

// InitDB func() error
func InitDB() error {
	var err error
	dbtype := "postgres"
	dbstr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", host, port, user, password, dbname)
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
