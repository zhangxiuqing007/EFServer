package controller

import (
	"EFServer/forum"
	"EFServer/tool"
	"EFServer/usecase"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var indexTemplate = template.Must(template.New("index").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/index.html"))))

type indexVM struct {
	loginInfo
	Themes []*forum.ThemeInDB
}

//Index 打开首页
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sendIndexPage(w, getExsitOrCreateNewSession(w, r, true))
}

//发送首页内容
func sendIndexPage(w http.ResponseWriter, s *Session) {
	vm := new(indexVM)
	vm.IsLogin = s.User != nil
	if vm.IsLogin {
		vm.UserName = s.User.Name
	}
	//把主题都放进去
	var err error
	vm.Themes, err = usecase.GetAllThemes()
	if err != nil {
		sendErrorPage(w, "查询主题列表失败")
		return
	} else if len(vm.Themes) == 0 {
		sendErrorPage(w, "无主题")
		return
	}
	indexTemplate.ExecuteTemplate(w, "index", vm)
}
