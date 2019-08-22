package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Index 打开首页
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sesstion := getExsitOrCreateNewSession(w, r)
	vm := new(indexVM)
	vm.IsLogin = sesstion.User != nil
	if vm.IsLogin {
		vm.UserName = sesstion.User.Name
	}
	indexTemplate.ExecuteTemplate(w, "index", vm)
}
