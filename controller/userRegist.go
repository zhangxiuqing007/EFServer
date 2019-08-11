package controller

import (
	"EFServer/forum"
	"EFServer/tool"
	"EFServer/usecase"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const modelNameOfRegist string = "regist"

var userRegistInputTemplate = template.Must(template.New(modelNameOfRegist).Parse(tool.ReadFileString("view/registInput.html")))
var userRegistSuccessTemplate = template.Must(template.New(modelNameOfRegist).Parse(tool.ReadFileString("view/registSuccess.html")))

type registInputVM struct {
	Error string
}

func readFormDataFromRegist(r *http.Request) (name string, account string, pwd1 string, pwd2 string) {
	strs := r.Form["name"]
	if len(strs) != 0 {
		name = strs[0]
	}
	strs = r.Form["account"]
	if len(strs) != 0 {
		account = strs[0]
	}
	strs = r.Form["password1"]
	if len(strs) != 0 {
		pwd1 = strs[0]
	}
	strs = r.Form["password2"]
	if len(strs) != 0 {
		pwd2 = strs[0]
	}
	return
}

//UserRegist UserRegist
func UserRegist(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		userRegistInputTemplate.ExecuteTemplate(w, modelNameOfRegist, registInputVM{Error: "请输入注册资料"})
		return
	}
	name, account, pwd1, pwd2 := readFormDataFromRegist(r)
	//如果是首次打开本页面
	if len(name) == 0 || len(account) == 0 || len(pwd1) == 0 || len(pwd2) == 0 {
		userRegistInputTemplate.ExecuteTemplate(w, modelNameOfRegist, registInputVM{Error: "请输入注册资料"})
		return
	}
	//后端再次检查一遍
	if pwd1 != pwd2 {
		userRegistInputTemplate.ExecuteTemplate(w, modelNameOfRegist, registInputVM{Error: "两次密码输入不一致"})
		return
	}
	//组织申请数据
	data := &usecase.UserSignUpData{
		Name:     name,
		Account:  account,
		Password: pwd1,
		UserType: forum.UserTypeNormalUser}
	//调用用例层代码，尝试添加账户，并返回错误
	err = usecase.AddUser(data)
	if err != nil {
		userRegistInputTemplate.ExecuteTemplate(w, modelNameOfRegist, registInputVM{Error: err.Error()})
		return
	}
	userRegistSuccessTemplate.ExecuteTemplate(w, modelNameOfRegist, nil)
}
