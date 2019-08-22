package controller

import (
	"net/http"

	"EFServer/usecase"
	"time"

	"github.com/julienschmidt/httprouter"
)

func readFormDataOfLogin(r *http.Request) (account string, pwd string) {
	strs := r.Form["account"]
	if strs != nil && len(strs) != 0 {
		account = strs[0]
	}
	strs = r.Form["password"]
	if strs != nil && len(strs) != 0 {
		pwd = strs[0]
	}
	return
}

//Login 登录页面
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	getExsitOrCreateNewSession(w, r).LastRequestTime = time.Now().UnixNano()
	loginInputTemplate.ExecuteTemplate(w, "login", &loginVM{Tip: "请输入账号密码"})
}

//LoginCommit 登录请求
func LoginCommit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		loginInputTemplate.ExecuteTemplate(w, "login", &loginVM{Tip: err.Error()})
		return
	}
	account, pwd := readFormDataOfLogin(r)
	//简单检查一下
	if len(account) == 0 || len(pwd) == 0 {
		loginInputTemplate.ExecuteTemplate(w, "login", &loginVM{Tip: "请输入账号密码"})
		return
	}
	//查询用户
	user, err := usecase.QueryUser(account, pwd)
	if err != nil {
		loginInputTemplate.ExecuteTemplate(w, "login", &loginVM{Tip: err.Error()})
		return
	}
	session := getExsitOrCreateNewSession(w, r)
	session.LastRequestTime = time.Now().UnixNano()
	session.User = user
	loginSuccessTemplate.ExecuteTemplate(w, "login", nil)
}

//Exit 登出
func Exit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	getExsitOrCreateNewSession(w, r).User = nil
	indexTemplate.ExecuteTemplate(w, "index", new(indexVM))
}
