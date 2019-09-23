package controller

import (
	"EFServer/forum"
	"EFServer/usecase"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var postTemplate = template.Must(template.ParseFiles("view/post.html", "view/comp/pageNavi.html", "view/comp/login.html"))

const cmtCountOnePage = 20                //帖子页，一页评论的数量
const halfPageCountToNavigationOfPost = 8 //评论导航页数量

type postVM struct {
	*loginInfo
	*forum.PostOnPostPage
	Comments []*forum.CmtOnPostPage
	*pageNavis
}

//PostInfo 查看帖子
func PostInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	postStrID := ps.ByName("postID")
	pageStrIndex := ps.ByName("pageIndex")
	//转化成int型的postId
	postID, err := strconv.ParseInt(postStrID, 10, 64)
	if err != nil {
		sendErrorPage(w, err.Error())
		return
	}
	pageIndex, err := strconv.ParseInt(pageStrIndex, 10, 64)
	if err != nil || pageIndex < 0 {
		pageIndex = 0
	}
	sendPostPage(w, postID, int(pageIndex), getExsitOrCreateNewSession(w, r, true))
}

//发送帖子页
func sendPostPage(w http.ResponseWriter, postID int64, pageIndex int, s *Session) {
	vm := new(postVM)
	vm.loginInfo = buildLoginInfo(s)
	//查询帖子主体信息
	var err error
	vm.PostOnPostPage, err = usecase.QueryPostOfPostPage(postID)
	if err != nil {
		sendErrorPage(w, "帖子查询失败")
	}
	//开始组织评论信息
	cmtTotalCount, err := usecase.QueryCommentsCountOfPost(postID)
	if err != nil {
		sendErrorPage(w, "查看帖子失败")
		return
	}
	//限制页Index
	pageIndex = limitPageIndex(pageIndex, cmtCountOnePage, cmtTotalCount)
	//查询评论内容
	vm.Comments, err = usecase.QueryCommentsOfPostPage(postID, cmtCountOnePage, pageIndex*cmtCountOnePage)
	if err != nil {
		sendErrorPage(w, "评论查询失败")
		return
	}
	//生成文字的日期 和评论所在的楼层
	baseLayerCount := pageIndex * cmtCountOnePage
	for i, v := range vm.Comments {
		v.FormatStringTime()
		v.FormatIndex(baseLayerCount + i)
	}
	//制作导航链接
	pathBuilder := func(index int) string {
		return fmt.Sprintf("/Post/%d/%d", postID, index)
	}
	beginIndex, endIndex := getNaviPageIndexs(pageIndex, cmtCountOnePage, halfPageCountToNavigationOfPost, cmtTotalCount)
	vm.pageNavis = buildPageNavis(pathBuilder, beginIndex, pageIndex, endIndex)
	//发送帖子页
	postTemplate.Execute(w, vm)
}
