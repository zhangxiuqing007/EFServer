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

	AddTheme(theme *forum.ThemeInDB) error            //新增主题
	DeleteTheme(themeID int) error                    //删除主题
	UpdateTheme(theme *forum.ThemeInDB) error         //更新主题（名称）
	QueryTheme(themeID int) (*forum.ThemeInDB, error) //查询主题
	QueryAllThemes() ([]*forum.ThemeInDB, error)      //查询所有主题
	QueryPostCountOfTheme(themeID int) (int, error)   //查询主题的帖子量

	AddPost(post *forum.PostInDB, cmt *forum.CommentInDB) error                                   //新增帖子
	AddPosts(post []*forum.PostInDB, cmt []*forum.CommentInDB) error                              //批量新增帖子
	DeletePost(postID int) error                                                                  //删除帖子
	QueryPost(postID int) (*forum.PostInDB, error)                                                //查询DB帖子
	QueryPostTitle(postID int) (string, error)                                                    //查询帖子标题
	UpdatePostTitle(*forum.PostInDB) error                                                        //修改帖子标题
	QueryPostsOfTheme(themeID int, count, offset, sortType int) ([]*forum.PostOnThemePage, error) //查询主题下的帖子列表
	QueryPostsOfUser(userID int, count, offset int) ([]*forum.PostOnThemePage, error)             //查询用户发的帖子列表
	QueryPostOfPostPage(postID int) (*forum.PostOnPostPage, error)                                //查询帖子页内，帖子的展示内容

	AddComment(comment *forum.CommentInDB) error                                                       //新增评论
	AddComments(comments []*forum.CommentInDB) error                                                   //批量增加评论
	DeleteComment(cmtID int) error                                                                     //删除评论
	QueryComment(cmtID int) (*forum.CommentInDB, error)                                                //查询DB评论
	UpdateComment(cmt *forum.CommentInDB) error                                                        //修改评论
	QueryComments(postID int) ([]*forum.CommentInDB, error)                                            //查询DB评论
	QueryCommentsOfPostPage(postID int, count, offset int, userID int) ([]*forum.CmtOnPostPage, error) //查询帖子页内，评论的展示内容

	AddPbItem(pb *forum.PBInDB) error                         //新增赞踩行
	QueryPbItem(cmtID int, userID int) (*forum.PBInDB, error) //查询赞踩行
	Praise(pb *forum.PBInDB) error                            //赞
	Belittle(pb *forum.PBInDB) error                          //贬
	PraiseCancel(pb *forum.PBInDB) error                      //取消赞
	BelittleCancel(pb *forum.PBInDB) error                    //取消贬
	SetPB(cmtID int, userID int, isP bool, isD bool) error    //尝试设置PB

	AddUser(user *forum.UserInDB) error                                                //新增用户
	DeleteUser(userID int) error                                                       //删除用户
	QueryUserByID(userID int) (*forum.UserInDB, error)                                 //通过id查询用户
	QueryUserByAccountAndPwd(account string, password string) (*forum.UserInDB, error) //通过账户密码查询用户
	QueryPostCountOfUser(userID int) (int, error)                                      //查询用户的发帖量
	IsUserNameExist(name string) bool                                                  //用户名是否存在
	IsUserAccountExist(account string) bool                                            //用户账号是否存在
}

const (
	//PostSortTypeCreatedTime 排序类型：发帖时间
	PostSortTypeCreatedTime = iota
	//PostSortTypeLastCmtTime 排序类型：最终评论时间
	PostSortTypeLastCmtTime
)
