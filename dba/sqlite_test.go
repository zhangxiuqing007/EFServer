package dba

import (
	"testing"
)

//连接至sqlite文件	go test -v -run Test_SqliteLink
func Test_SqliteLink(t *testing.T) {
	db := new(SqliteIns)
	err := db.Open("../ef.db")
	if err != nil {
		_, err = db.QueryAllThemes()
	}
	if err != nil {
		t.Error("x失败：连接sqlite文件")
		t.FailNow()
	} else {
		t.Logf("成功：连接至sqlite文件")
	}
	db.Close()
}
