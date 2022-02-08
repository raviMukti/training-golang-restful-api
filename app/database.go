package app

import (
	"database/sql"
	"time"

	"github.com/raviMukti/training-golaang-restful-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/training_db?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
