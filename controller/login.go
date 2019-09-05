package controller

import (
	"html/template"
	"net/http"

	"EFServer/tool"
	"EFServer/usecase"

	"github.com/julienschmidt/httprouter"
)

var loginInputTemplate = template.Must(template.New("login").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/loginInput.html"))))
var loginSuccessTemplate = template.Must(template.New("login").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/loginSuccess.html"))))

type loginVM struct {
	Tip string
}

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
	session := getExsitOrCreateNewSession(w, r, true)
	session.User = user
	loginSuccessTemplate.ExecuteTemplate(w, "login", user.Name)
}

//Exit 登出
func Exit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	s := getExsitOrCreateNewSession(w, r, true)
	s.User = nil
	sendIndexPage(w, s)
}
