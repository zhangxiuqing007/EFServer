package controller

import (
	"EFServer/forum"
	"EFServer/tool"
	"EFServer/usecase"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var userTemplate = template.Must(template.New("user").Parse(tool.MustStr(tool.ReadAllTextUtf8("view/user.html"))))

type userVM struct {
	ID                                                                    int64
	Name, SignUpTime, Type, State, LastOperateTime                        string
	PostTotalCount, CmtTotalCount, TotalPraisedTimes, TotalBelittledTimes int
}

//UserInfo 查看用户个人资料
func UserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userIDStr := ps.ByName("userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		sendErrorPage(w, err.Error())
		return
	}
	sendUserPage(w, userID, getExsitOrCreateNewSession(w, r, true))
}

//统计并发送用户资料页面
func sendUserPage(w http.ResponseWriter, userID int64, s *Session) {
	//db统计用户信息
	saInfo, err := usecase.QueryUserSaInfo(userID)
	if err != nil {
		sendErrorPage(w, err.Error())
		return
	}
	vm := new(userVM)
	vm.ID = saInfo.ID
	vm.Name = saInfo.Name
	vm.SignUpTime = tool.FormatTimeDetail(time.Unix(0, saInfo.SignUpTime))
	vm.Type = forum.GetUserTypeShowName(saInfo.UserType)
	vm.State = forum.GetUserStateShowName(saInfo.UserState)
	vm.LastOperateTime = tool.FormatTimeDetail(time.Unix(0, saInfo.LastOperateTime))
	vm.PostTotalCount = saInfo.PostTotalCount
	vm.CmtTotalCount = saInfo.CmtTotalCount
	vm.TotalPraisedTimes = saInfo.TotalPraisedTimes
	vm.TotalBelittledTimes = saInfo.TotalBelittledTimes
	userTemplate.ExecuteTemplate(w, "user", vm)
}
