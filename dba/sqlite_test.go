package dba

import "testing"
import "EFServer/forum"
import "time"

func Test_Sqlite_AddUser(t *testing.T) {
	ioTool := &SqliteIns{}
	ioTool.Open("../ef.db")
	defer ioTool.Close()

	newUser := new(forum.User)
	newUser.Name = "我最大最二"
	newUser.Account = "690313521_2@qq.com"
	newUser.PassWord = "aa135828"
	newUser.SignUpTime = time.Now().UnixNano()
	newUser.UserState = forum.UserStateNormal
	newUser.UserType = forum.UserTypeNormalUser

	err := ioTool.AddUser(newUser)
	if err != nil {
		t.Error(err)
	}
}
