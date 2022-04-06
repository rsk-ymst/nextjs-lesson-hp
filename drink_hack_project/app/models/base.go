package models

import (
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	USER_NAME = "users"
	QUERY_USER_TABLE =
` CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`
)

func init() {
	Db, err = sql.Open("sqlite", config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(QUERY_USER_TABLE, USER_NAME)

	Db.Exec(cmdU)
}