package controller

import (
	"EFServer/forum"
	"EFServer/tool"
	"EFServer/usecase"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var postTemplate = template.Must(template.New("post").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/post.html"))))

type postVM struct {
	loginInfo
	*forum.PostOnPostPage
	Comments []*forum.CmtOnPostPage
}

//PostInfo 查看帖子
func PostInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	postStrID := ps.ByName("postID")
	//转化成int型的postId
	postID, err := strconv.ParseInt(postStrID, 10, 64)
	if err != nil {
		sendErrorPage(w, err.Error())
		return
	}
	sendPostPage(w, postID, getExsitOrCreateNewSession(w, r, true))
}

//发送帖子页
func sendPostPage(w http.ResponseWriter, postID int64, s *Session) {
	vm := new(postVM)
	vm.IsLogin = s.User != nil
	if vm.IsLogin {
		vm.UserName = s.User.Name
	}
	//查询帖子信息
	pgPost, err := usecase.QueryPostPG(postID)
	if err != nil {
		sendErrorPage(w, "帖子查询失败")
	}
	vm.PostOnPostPage = pgPost
	//查询评论内容
	vm.Comments, err = usecase.QueryPgComments(postID)
	if err != nil {
		sendErrorPage(w, "评论查询失败")
	}
	//生成文字的日期 和评论所在的楼层
	for i, v := range vm.Comments {
		v.FormatStringTime()
		v.FormatIndex(i)
	}
	//发送帖子页
	postTemplate.ExecuteTemplate(w, "post", vm)
}
