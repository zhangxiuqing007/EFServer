package controller

import (
	"EFServer/forum"
	"EFServer/tool"
	"EFServer/usecase"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var themeTemplate = template.Must(template.New("theme").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/theme.html"))))

type themeVM struct {
	ThemeTitle string
	IsLogin    bool
	UserName   string
	Posts      []*forum.Post
}

//Theme 进入主题
func Theme(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	themeName := ps.ByName("theme")
	if themeName == "" {
		sendErrorPage(w, "无此主题")
		return
	}
	sendThemePage(w, themeName, getExsitOrCreateNewSession(w, r, true))
}

func sendThemePage(w http.ResponseWriter, themeName string, s *Session) {
	thm, err := usecase.GetTheme(themeName)
	if err != nil {
		sendErrorPage(w, "访问主题失败")
		return
	}
	vm := new(themeVM)
	vm.ThemeTitle = themeName
	vm.IsLogin = s.User != nil
	if vm.IsLogin {
		vm.UserName = s.User.Name
	}
	//获取帖子列表

	themeTemplate.ExecuteTemplate(w, "theme", vm)
}
