package main

import (
	"EFServer/dba"
	"EFServer/forum"
	"fmt"
)

func main() {
	fmt.Println("hello 世界")
	//给接口赋值
	user := &forum.User{
		ID:       0,
		UserCode: "zxq",
		PassWord: "asd   ",
		UserType: forum.Administrator,
	}
	dba.DataOper.AddUser(user)
	dba.DataOper.DeleteUser(user)
}
