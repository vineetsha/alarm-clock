package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"sync"
)

var DatabaseConnection *sql.DB
var once sync.Once

func ConnectDB() {
	once.Do(func() {
		cfg := mysql.NewConfig()
		cfg.User = "root"
		cfg.Passwd = ""
		cfg.Net = "tcp"
		cfg.Addr = "127.0.0.1:3306"
		cfg.DBName = "getgoing"
		var err error
		DatabaseConnection, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			panic(err)
		} else {
			fmt.Println("DB connected")
		}
	})
}

func PingDB() error {
	if DatabaseConnection != nil {
		return DatabaseConnection.Ping()
	}
	return errors.New("db not initialized")
}

func ShutDownDB() {
	if DatabaseConnection != nil {
		fmt.Println("Closing DB connection")
		DatabaseConnection.Close()
	}
}
