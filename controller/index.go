package controller

import (
	"net/http"

	"EFServer/tool"
	"html/template"

	"github.com/julienschmidt/httprouter"
)

var indexTemplate = template.Must(template.New("index").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/index.html"))))

type indexVM struct {
	Login bool
	Name  string
}

//Index firstPage
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sesstion := getExsitOrCreateNewSession(w, r)
	vm := new(indexVM)
	vm.Login = sesstion.User != nil
	if vm.Login {
		vm.Name = sesstion.User.Name
	}
	indexTemplate.ExecuteTemplate(w, "index", vm)
}

//Exit 登出
func Exit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	getExsitOrCreateNewSession(w, r).User = nil
	indexTemplate.ExecuteTemplate(w, "index", new(indexVM))
}
