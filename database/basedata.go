package database

import (
	"database/sql"
	"ginlearning/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:0061729940@tcp(localhost:3306)/ginbase")
	helper.PanicErrorIf(err)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxOpenConns(20)
	return db
}
