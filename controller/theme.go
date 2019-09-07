package controller

import (
	"EFServer/forum"
	"EFServer/tool"
	"EFServer/usecase"
	"html/template"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var themeTemplate = template.Must(template.New("theme").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/theme.html"))))

type themeVM struct {
	ThemeTitle  string
	IsLogin     bool
	UserName    string
	PostHeaders []*postHeader
}

type postHeader struct {
	*forum.PostBriefInfo
	FcreatedTime string
	FlastCmtTime string
}

//Theme 进入主题
func Theme(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	themeName := ps.ByName("theme")
	if themeName == "" {
		sendErrorPage(w, "无目标主题")
		return
	}
	sendThemePage(w, themeName, getExsitOrCreateNewSession(w, r, true))
}

//发送主题页，帖子列表
func sendThemePage(w http.ResponseWriter, themeName string, s *Session) {
	tm, err := usecase.GetTheme(themeName)
	if err != nil {
		sendErrorPage(w, "访问主题失败"+err.Error())
		return
	}
	vm := new(themeVM)
	vm.ThemeTitle = "边缘社区-" + themeName
	vm.IsLogin = s.User != nil
	if vm.IsLogin {
		vm.UserName = s.User.Name
	}
	//获取帖子列表
	tempHeaders, err := usecase.QueryPosts(tm.ID, 20, 0, 0)
	if err != nil {
		sendErrorPage(w, "查询帖子列表失败")
		return
	}
	vmHeaders := make([]*postHeader, 0, 20)
	formatTime := func(ticks int64) string {
		return tool.FormatTimeDetail(time.Unix(0, ticks))
	}
	for _, v := range tempHeaders {
		vmHeaders = append(vmHeaders, &postHeader{v, formatTime(v.CreateTime), formatTime(v.LastCmtTime)})
	}
	vm.PostHeaders = vmHeaders
	themeTemplate.ExecuteTemplate(w, "theme", vm)
}
