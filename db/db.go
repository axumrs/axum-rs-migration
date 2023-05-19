package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	// PostgreSQL 连接
	PG *sqlx.DB
	// MySQL 连接
	MY *sqlx.DB
)

func InitPG(dsn string) (err error) {
	PG, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return err
	}

	return nil
}

func InitMySQL(dsn string) (err error) {
	MY, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}
