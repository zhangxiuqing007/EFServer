package controller

import (
	"net/http"

	"EFServer/forum"
	"EFServer/tool"
	"html/template"

	"github.com/julienschmidt/httprouter"
)

var indexTemplate = template.Must(template.New("index").Parse(tool.ReadFileString("view/index.html")))

type indexVM struct {
	Login bool
	User  *forum.User
}

//Index firstPage
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sesstion := getExsitOrCreateNewSession(w, r)
	vm := new(indexVM)
	vm.Login = sesstion.User != nil
	vm.User = sesstion.User
	indexTemplate.ExecuteTemplate(w, "index", vm)
}
