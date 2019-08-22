package dba

import (
	"database/sql"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func linkToMysql() (*sql.DB, error) {
	return sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/efdb?charset=utf8")
}
