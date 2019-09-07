package usecase

import "EFServer/forum"

var db IDataIO

//SetDbInstance 设置当前的db实现
func SetDbInstance(dbIns IDataIO) {
	db = dbIns
}

//IDataIO IO接口
type IDataIO interface {
	Open(string) error
	Close() error

	AddTheme(themeName string) error
	DeleteTheme(themeID int64) error
	UpdateTheme(theme *forum.Theme) error
	QueryTheme(themeName string) (*forum.Theme, error)
	QueryThemes() ([]*forum.Theme, error) //查询所有主题

	AddPost(post *forum.Post) error
	DeletePost(postID int64) error
	UpdatePost(post *forum.Post) error
	QueryPosts(themeID int64, count, offset, sortType int) ([]*forum.PostBriefInfo, error) //查询主题下的所有帖子简要内容
	QueryPost(postID int64) (*forum.Post, error)                                           //comments有内容

	AddComment(comment *forum.Comment) error
	DeleteComment(cmtID int64) error
	UpdateComment(comment *forum.Comment) error
	QueryComments(postID int64) ([]*forum.Comment, error)

	AddUser(user *forum.User) error
	DeleteUser(userID int64) error
	UpdateUser(user *forum.User) error
	QueryUserByID(userID int64) (*forum.User, error)
	QueryUserByAccountAndPwd(account string, password string) (*forum.User, error)

	IsUserNameExist(name string) bool
	IsUserAccountExist(account string) bool
}

const (
	//PostSortTypeCreatedTime 排序类型：发帖时间
	PostSortTypeCreatedTime = iota
	//PostSortTypeLastCmtTime 排序类型：最终评论时间
	PostSortTypeLastCmtTime
)
