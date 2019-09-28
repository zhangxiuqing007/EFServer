package dba

import (
	"testing"
)

//测试查询所有主题的速度
func BenchmarkReadAllThemes(b *testing.B) {
	rander := new(testResourceBuilder)
	sqlIns := rander.buildCurrentTestSQLIns()
	for i := 0; i < b.N; i++ {
		_, _ = sqlIns.QueryAllThemes()
	}
}
