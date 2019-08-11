package controller

import (
	"net/http"

	"EFServer/tool"
	"EFServer/usecase"
	"html/template"

	"github.com/julienschmidt/httprouter"
)

const modelNameOfLogin string = "login"

var loginInputTemplate = template.Must(template.New(modelNameOfLogin).Parse(tool.ReadFileString("view/loginInput.html")))
var loginSuccessTemplate = template.Must(template.New(modelNameOfLogin).Parse(tool.ReadFileString("view/loginSuccess.html")))

type loginVM struct {
	Error string
}

func readFormDataOfLogin(r *http.Request) (account string, pwd string) {
	strs := r.Form["account"]
	if len(strs) != 0 {
		account = strs[0]
	}
	strs = r.Form["password"]
	if len(strs) != 0 {
		pwd = strs[0]
	}
	return
}

//Login Login
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		loginInputTemplate.ExecuteTemplate(w, modelNameOfLogin, &loginVM{Error: err.Error()})
		return
	}
	account, pwd := readFormDataOfLogin(r)
	if len(account) == 0 || len(pwd) == 0 {
		loginInputTemplate.ExecuteTemplate(w, modelNameOfLogin, nil)
		return
	}
	user, err := usecase.QueryUser(account, pwd)
	if err != nil {
		loginInputTemplate.ExecuteTemplate(w, modelNameOfLogin, &loginVM{Error: err.Error()})
		return
	}
	getExsitOrCreateNewSession(w, r).User = user
	loginSuccessTemplate.ExecuteTemplate(w, modelNameOfLogin, nil)
}
