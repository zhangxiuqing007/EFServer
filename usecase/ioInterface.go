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
	Clear() error
	Close() error

	AddTheme(theme *forum.ThemeInDB) error              //新增主题
	DeleteTheme(themeID int64) error                    //删除主题
	UpdateTheme(theme *forum.ThemeInDB) error           //更新主题（名称）
	QueryTheme(themeID int64) (*forum.ThemeInDB, error) //查询主题
	QueryAllThemes() ([]*forum.ThemeInDB, error)        //查询所有主题

	AddPost(post *forum.PostInDB) error                                                             //新增帖子
	AddPosts(post []*forum.PostInDB) error                                                          //批量新增帖子
	DeletePost(postID int64) error                                                                  //删除帖子
	QueryPost(postID int64) (*forum.PostInDB, error)                                                //查询DB帖子
	QueryPostCountOfTheme(themeID int64) (int, error)                                               //查询主题帖子总数量
	QueryPostCountOfUser(userID int64) (int, error)                                                 //查询用户发帖总数量
	QueryPostsOfTheme(themeID int64, count, offset, sortType int) ([]*forum.PostOnThemePage, error) //查询主题下的帖子列表
	QueryPostsOfUser(userID int64, count, offset int) ([]*forum.PostOnThemePage, error)             //查询用户发的帖子列表
	QueryPostOfPostPage(postID int64) (*forum.PostOnPostPage, error)                                //查询帖子页内，帖子的展示内容

	AddComment(comment *forum.CommentInDB) error                                             //新增评论
	AddComments(comments []*forum.CommentInDB) error                                         //批量增加评论
	DeleteComment(cmtID int64) error                                                         //删除评论
	QueryComments(postID int64) ([]*forum.CommentInDB, error)                                //查询DB评论
	QueryCommentsCountOfPost(postID int64) (int, error)                                      //统计帖子的评论数量
	QueryCommentsOfPostPage(postID int64, count, offset int) ([]*forum.CmtOnPostPage, error) //查询帖子页内，评论的展示内容

	AddUser(user *forum.UserInDB) error                                                //新增用户
	DeleteUser(userID int64) error                                                     //删除用户
	QueryUserByID(userID int64) (*forum.UserInDB, error)                               //通过id查询用户
	QueryUserByAccountAndPwd(account string, password string) (*forum.UserInDB, error) //通过账户密码查询用户
	QueryUserSaInfoByID(userID int64) (*forum.UserStatisticsInfo, error)               //查询用户的统计信息
	IsUserNameExist(name string) bool                                                  //用户名是否存在
	IsUserAccountExist(account string) bool                                            //用户账号是否存在
}

const (
	//PostSortTypeCreatedTime 排序类型：发帖时间
	PostSortTypeCreatedTime = iota
	//PostSortTypeLastCmtTime 排序类型：最终评论时间
	PostSortTypeLastCmtTime
)
