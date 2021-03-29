package helpers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", GetEnv("DBUSERNAME"), GetEnv("DBPASSWORD"), GetEnv("DBHOST"), GetEnv("DBPORT"), GetEnv("DBDATABASE"))

func ConnectMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	LogFatal(err)
	return db, nil
}
