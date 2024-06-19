package handler

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
)

const (
	username = "root"
	password = "safesync"
	host     = "10.1.103.111"
	port     = 3306
	dbName   = "central"
)

var dbx *MysqlDB

// var db *sql.DB
var err error
var once sync.Once

type MysqlDB struct {
	db *sqlx.DB
}

func connectDB() (*MysqlDB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbName))
	if err != nil {
		err = fmt.Errorf("sql.Open fail, err(%v)", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return &MysqlDB{db}, nil
}

func (m *MysqlDB) Close() {
	m.db.Close()
}

// func connectDB() (*sql.DB, error) {
// 	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbName))
// 	if err != nil {
// 		err = fmt.Errorf("sql.Open fail, err(%v)", err)
// 		return nil, err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

func initDB() error {
	dbx, err = connectDB()
	if err != nil {
		err = fmt.Errorf("connectDB fail, err(%v)", err)
		return err
	}

	return nil
}

func getDB() (*MysqlDB, error) {
	once.Do(func() {
		err = initDB()
		if err != nil {
			err = fmt.Errorf("initDB fail, err(%w)", err)
			return
		}
	})

	return dbx, err
}
