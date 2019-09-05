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
	IsLogin  bool
	UserName string
	Themes   []*forum.Theme
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
	buildEmptyThemes := func(content string) []*forum.Theme {
		tms := make([]*forum.Theme, 0, 1)
		tms = append(tms, &forum.Theme{ID: -1, Name: content})
		return tms
	}
	if err != nil {
		vm.Themes = buildEmptyThemes("读取主题列表失败")
	} else if len(vm.Themes) == 0 {
		vm.Themes = buildEmptyThemes("无主题")
	}
	indexTemplate.ExecuteTemplate(w, "index", vm)
}
