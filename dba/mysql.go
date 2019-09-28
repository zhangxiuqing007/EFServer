package dba

import (
	"database/sql"
	"fmt"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//MySQLIns mysql数据库实现
type MySQLIns struct {
	sqlBase
}

//Open 打开
func (s *MySQLIns) Open(dbFilePath string) error {
	var err error
	s.db, err = sql.Open("mysql", fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/efdb", dbFilePath))
	return err
}

//Clear 清空
func (s *MySQLIns) Clear() error {
	const sqlStrToClear = `
	delete from cmt;
	delete from post;
	delete from theme;
	delete from user;`
	_, err := s.db.Exec(sqlStrToClear)
	return err
}
