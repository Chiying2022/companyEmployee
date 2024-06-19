package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func InitMysqlDB(user, pwd, url, dbname string, port, maxConns int) (*MysqlDB, error) {
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local",
		user, pwd, url, port, dbname,
	)

	var db *sqlx.DB
	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "建立 db connection 錯誤")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "mysql ping 錯誤")
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(maxConns)

	return &MysqlDB{db}, err
}

func (m *MysqlDB) Close() {
	m.db.Close()
}

type MysqlDB struct {
	db *sqlx.DB
}
