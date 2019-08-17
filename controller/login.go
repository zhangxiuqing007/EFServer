package controller

import (
	"net/http"

	"EFServer/tool"
	"EFServer/usecase"
	"html/template"

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

//Login Login
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	loginInputTemplate.ExecuteTemplate(w, "login", &loginVM{Tip: "请输入账号密码"})
}

//LoginCommit LoginCommit
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
	user, err := usecase.QueryUser(account, pwd)
	if err != nil {
		loginInputTemplate.ExecuteTemplate(w, "login", &loginVM{Tip: err.Error()})
		return
	}
	getExsitOrCreateNewSession(w, r).User = user
	loginSuccessTemplate.ExecuteTemplate(w, "login", nil)
}
