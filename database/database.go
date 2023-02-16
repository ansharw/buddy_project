package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/buddyku?parseTime=true")
	if err != nil {
		panic(err)
	}

	// db.SetMaxIdleConns(10)
	// db.SetMaxOpenConns(100)
	// db.SetConnMaxIdleTime(30 * time.Second)
	// db.SetConnMaxLifetime(12 * time.Minute)
	return db
}
