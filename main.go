package main

import (
	"EFServer/controller"
	"EFServer/dba"
	"EFServer/usecase"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("启动程序...")
	//设置db实现
	usecase.SetDbInstance(new(dba.Abs))
	//非表单路由
	router := httprouter.New()
	router.GET("/", controller.Index)
	router.GET("/UserRegist", controller.UserRegist)
	router.GET("/UserRegistCommit", controller.UserRegistCommit)
	router.GET("/Login", controller.Login)
	router.GET("/LoginCommit", controller.LoginCommit)
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		fmt.Print("程序启动失败：" + err.Error())
		os.Exit(0)
	}
}
