package main

import (
	"EFServer/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("service is starting...")
	//路由，非表单
	router := httprouter.New()
	router.GET("/", controller.Index)
	router.GET("/UserRegist", controller.UserRegist)
	router.GET("/Login", controller.Login)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
