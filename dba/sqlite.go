package dba

import (
	"database/sql"

	//sqlite3 driver
	_ "github.com/mattn/sqlite3"
)

//SqliteIns sqlite实现
type SqliteIns struct {
	sqlBase
}

//Open 打开
func (s *SqliteIns) Open(dbFilePath string) error {
	var err error
	s.db, err = sql.Open("sqlite3", dbFilePath)
	return err
}
