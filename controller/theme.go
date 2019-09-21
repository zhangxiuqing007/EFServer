package controller

import (
	"EFServer/forum"
	"EFServer/tool"
	"EFServer/usecase"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var themeTemplate = template.Must(template.New("theme").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/theme.html"))))

const postCountOnePage = 20          //主题页，一页帖子数量
const halfPageCountToNavigation = 10 //导航页数量

type themeVM struct {
	WebTitle  string //网页Header
	ThemeName string //主题名

	loginInfo //登录信息

	PostHeaders []*forum.PostOnThemePage //帖子简要内容

	HeadPageNavis []*pageNaviVM //前导航页
	CurrentPage   *pageNaviVM   //当前页
	TailPageNavis []*pageNaviVM //后导航页
}

type pageNaviVM struct {
	Path   string
	Number int
}

//Theme 请求主题内帖子列表
func Theme(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	themeID, _ := strconv.ParseInt(ps.ByName("id"), 10, 64)
	pageIndex, _ := strconv.ParseInt(ps.ByName("page"), 10, 64)
	//检查
	if pageIndex < 0 {
		pageIndex = 0
	}
	sendThemePage(w, themeID, int(pageIndex), getExsitOrCreateNewSession(w, r, true))
}

//发送主题页，帖子列表
func sendThemePage(w http.ResponseWriter, themeID int64, pageIndex int, s *Session) {
	tm, err := usecase.GetTheme(themeID)
	if err != nil {
		sendErrorPage(w, "访问主题失败")
		return
	}
	//创建 viewModel对象
	vm := new(themeVM)
	//给vm赋基本值
	vm.WebTitle = "边缘社区-" + tm.Name
	vm.ThemeName = tm.Name
	vm.IsLogin = s.User != nil
	if vm.IsLogin {
		vm.UserName = s.User.Name
	}
	//获取帖子列表，根据请求的页码查询帖子列表
	vm.PostHeaders, err = usecase.QueryPosts(tm.ID, postCountOnePage, pageIndex*postCountOnePage, s.PostSortType)
	if err != nil {
		sendErrorPage(w, "查询帖子列表失败")
		return
	}
	for _, v := range vm.PostHeaders {
		v.FormatStringTime()
	}
	//path制作函数
	buildPath := func(i int) string {
		return fmt.Sprintf("/Theme/%d/%d", tm.ID, i)
	}
	totalPostCount := usecase.QueryPostCountOfTheme(tm.ID)
	//确定导航页限制
	beginIndex, endIndex := getNaviPageIndexs(pageIndex, postCountOnePage, halfPageCountToNavigation, totalPostCount)
	//制作前序导航页
	vm.HeadPageNavis = make([]*pageNaviVM, 0, halfPageCountToNavigation)
	for i := beginIndex; i < pageIndex; i++ {
		vm.HeadPageNavis = append(vm.HeadPageNavis, &pageNaviVM{buildPath(i), i + 1})
	}
	//设定当前页
	vm.CurrentPage = &pageNaviVM{"", pageIndex + 1}
	//制作后续导航页
	vm.TailPageNavis = make([]*pageNaviVM, 0, halfPageCountToNavigation)
	for i := pageIndex + 1; i <= endIndex; i++ {
		vm.TailPageNavis = append(vm.TailPageNavis, &pageNaviVM{buildPath(i), i + 1})
	}
	themeTemplate.ExecuteTemplate(w, "theme", vm)
}
