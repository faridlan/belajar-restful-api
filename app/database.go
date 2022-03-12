package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/faridlan/belajar-restful-api/helper"
)

func NewDB() *sql.DB {
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang")
	db, err := sql.Open("mysql", fmt.Sprintf("root:root@tcp(%s:%s)/belajar_golang", host, port))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
