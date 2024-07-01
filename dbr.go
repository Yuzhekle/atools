package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr/v2"
)

func newDBConnection() *dbr.Connection {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "123456", "127.0.0.1:3380", "tms")
	if dbConnection, openErr := dbr.Open("mysql", dsn, nil); openErr != nil {
		panic(openErr)
	} else {
		dbConnection.SetMaxIdleConns(2)
		dbConnection.SetMaxOpenConns(10)
		dbConnection.SetConnMaxLifetime(6 * time.Second)
		return dbConnection
	}
}

func main() {
	db := newDBConnection()
	fmt.Println(db)
	sess := db.NewSession(nil)
	fmt.Println(sess)

	var id int64
	sess.Select("instant_room_ct").From("ebk_daily_rate").Where("daily_rate_id = ?", 12454428).Load(&id)
	fmt.Println(id)

	res, err := sess.Update("ebk_daily_rate").
		Set("instant_room_ct", 12).
		Set("status", 3).
		Set("update_date", time.Now()).
		Where("daily_rate_id = ?", 12454428).Exec()
	if err != nil {
		panic(err)
	}
	rows, _ := res.RowsAffected()
	fmt.Println(rows)
}

func sqlx() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3380))")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(6 * time.Second)
	fmt.Println(db)
}
