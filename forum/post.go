package forum

import (
	"EFServer/tool"
	"strconv"
	"time"
)

//PostOnPostPage 帖子，在帖子页展示的内容
type PostOnPostPage struct {
	ID    int64
	Title string

	ThemeID   int64
	ThemeName string
}

//CmtOnPostPage 评论在帖子页展示的信息
type CmtOnPostPage struct {
	ID       int64
	IndexStr string
	Content  string

	PraiseTimes   int
	BelittleTimes int

	CmterID   int64
	CmterName string
	CmtTime   int64
	CmtTimeF  string
}

//FormatStringTime 生成文字类型的时间
func (cmt *CmtOnPostPage) FormatStringTime() {
	cmt.CmtTimeF = tool.FormatTimeDetail(time.Unix(0, cmt.CmtTime))
}

//FormatIndex 生成楼层字符
func (cmt *CmtOnPostPage) FormatIndex(index int) {
	if index == 0 {
		cmt.IndexStr = "楼主"
	} else {
		cmt.IndexStr = strconv.Itoa(index) + "楼"
	}
}

//PostOnThemePage 帖子在主题页中展示的信息
type PostOnThemePage struct {
	ID    int64
	Title string

	CmtCount int

	CreaterID    int64
	CreaterName  string
	CreateTime   int64
	CreatedTimeF string

	LastCmterID   int64
	LastCmterName string
	LastCmtTime   int64
	LastCmtTimeF  string
}

//FormatStringTime 生成文字类型的时间
func (p *PostOnThemePage) FormatStringTime() {
	p.CreatedTimeF = tool.FormatTimeDetail(time.Unix(0, p.CreateTime))
	p.LastCmtTimeF = tool.FormatTimeDetail(time.Unix(0, p.LastCmtTime))
}

//FixCmtCount 修正评论数量，减去主楼
func (p *PostOnThemePage) FixCmtCount() {
	p.CmtCount--
}

//PostInDB 帖子，数据库形态
type PostInDB struct {
	ID      int64
	ThemeID int64
	UserID  int64

	Title string

	State int
}

//CommentInDB 评论，数据库形态
type CommentInDB struct {
	ID     int64
	PostID int64
	UserID int64

	Content string

	State int

	CreatedTime  int64
	LastEditTime int64
	EditTimes    int

	PraiseTimes   int
	BelittleTimes int
}

const (
	//PostStateNormal 帖子状态：正常
	PostStateNormal = iota
	//PostStateLock 帖子状态：锁定
	PostStateLock
	//PostStateHide 帖子状态：隐藏
	PostStateHide
)

const (
	//CmtStateNormal 评论状态：正常
	CmtStateNormal = iota
	//CmtStateLock 评论状态：锁定
	CmtStateLock
	//CmtStateHide 评论状态：隐藏
	CmtStateHide
)
