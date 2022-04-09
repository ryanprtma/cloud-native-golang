package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = "root"
	database string = "models"
	host     string = "tcp(app-mysql:3308)"
	// host string = "tcp(127.0.0.1:3306)"
)

var (
	dsn = fmt.Sprintf("%v:%v@%v/%v", username, password, host, database)
)

// HubToMySQL
func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
