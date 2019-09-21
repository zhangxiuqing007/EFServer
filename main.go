package main

import (
	"EFServer/controller"
	"EFServer/dba"
	"EFServer/usecase"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("启动程序...")
	//db实现
	sqlIns := new(dba.SqliteIns)
	err := sqlIns.Open("ef.db")
	if err != nil {
		panic(err)
	}
	defer sqlIns.Close()
	usecase.SetDbInstance(sqlIns)

	//URL路由
	router := httprouter.New()
	router.GET("/", controller.Index)
	router.GET("/UserRegist", controller.UserRegist)
	router.GET("/UserRegistCommit", controller.UserRegistCommit)
	router.GET("/Login", controller.Login)
	router.GET("/LoginCommit", controller.LoginCommit)
	router.GET("/Exit", controller.Exit)
	router.GET("/Theme/:id/:page", controller.Theme)
	router.GET("/User/:userID", controller.UserInfo)
	router.GET("/Post/:postID", controller.PostInfo)
	fmt.Println("开始监听HTTP请求...")
	err = http.ListenAndServe("localhost:15856", router)
	if err != nil {
		fmt.Print("程序启动失败：" + err.Error())
	}
}
