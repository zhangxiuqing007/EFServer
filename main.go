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
	fmt.Println("正在启动程序...")
	const mysql = true
	var sqlIns usecase.IDataIO
	if mysql {
		sqlIns = new(dba.MySQLIns)
		if err := sqlIns.Open("mysql5856"); err != nil {
			panic(err)
		}
	} else {
		//db实现
		sqlIns = new(dba.SqliteIns)
		if err := sqlIns.Open("ef.db"); err != nil {
			panic(err)
		}
	}
	defer sqlIns.Close()
	usecase.SetDbInstance(sqlIns)
	//URL路由
	router := httprouter.New()
	router.GET("/", controller.Index)

	router.GET("/UserRegist", controller.UserRegist)
	router.POST("/UserRegistCommit", controller.UserRegistCommit)

	router.GET("/Login", controller.Login)
	router.POST("/LoginCommit", controller.LoginCommit)

	router.GET("/Exit", controller.Exit)

	router.GET("/Theme/:themeID/:pageIndex", controller.Theme)

	router.GET("/User/:userID", controller.UserInfo)
	router.GET("/User/:userID/:pageIndex", controller.UserPosts)

	router.GET("/Post/Content/:postID/:pageIndex", controller.PostInfo)
	router.GET("/Post/TitleEdit/:postID", controller.PostTitleEdit)
	router.POST("/Post/TitleEditSubmit", controller.PostTitleEditCommit)

	router.GET("/NewPostInput/:themeID", controller.NewPostInput)
	router.POST("/NewPostCommit", controller.NewPostCommit)

	router.POST("/Cmt", controller.Cmt)
	router.GET("/Cmt/Edit/:cmtID/:cmtPageIndex", controller.CmtEdit)
	router.POST("/Cmt/EditSubmit", controller.CmtEditCommit)
	router.POST("/Cmt/PG", controller.CmtPb)

	fmt.Println("开始监听HTTP请求...")
	if err := http.ListenAndServe("localhost:15856", router); err != nil {
		fmt.Print("程序启动失败：" + err.Error())
	}
}
