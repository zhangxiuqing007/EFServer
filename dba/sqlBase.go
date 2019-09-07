package dba

import (
	"EFServer/forum"
	"EFServer/tool"
	"database/sql"
)

type sqlBase struct {
	db *sql.DB
}

//Close 关闭
func (s *sqlBase) Close() error {
	return s.db.Close()
}

const sqlStrToAddTheme = "insert into theme (name) values (?)"

//AddTheme 增加主题
func (s *sqlBase) AddTheme(themeName string) error {
	_, err := s.db.Exec(sqlStrToAddTheme, themeName)
	return err
}

const sqlStrToDeleteTheme = "delete from theme where ID = ?"

//DeleteTheme 删除主题
func (s *sqlBase) DeleteTheme(themeID int64) error {
	_, err := s.db.Exec(sqlStrToDeleteTheme, themeID)
	return err
}

const sqlStrToUpdateTheme = "update theme set name = ? where ID = ?"

//UpdateTheme 更新主题
func (s *sqlBase) UpdateTheme(theme *forum.Theme) error {
	_, err := s.db.Exec(sqlStrToUpdateTheme, theme.Name, theme.ID)
	return err
}

const sqlStrToQueryTheme = "select * from theme where name = ?"

//QueryTheme 查询某个主题
func (s *sqlBase) QueryTheme(themeName string) (*forum.Theme, error) {
	tm := new(forum.Theme)
	err := s.db.QueryRow(sqlStrToQueryTheme, themeName).Scan(&tm.ID, &tm.Name)
	return tm, err
}

const sqlStrToQueryThemes = "select * from theme"

//QueryThemes 查询主题列表
func (s *sqlBase) QueryThemes() ([]*forum.Theme, error) {
	rows, err := s.db.Query(sqlStrToQueryThemes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tms := make([]*forum.Theme, 0, 12)
	for rows.Next() {
		tm := new(forum.Theme)
		err = rows.Scan(&tm.ID, &tm.Name)
		if err != nil {
			return nil, err
		}
		tms = append(tms, tm)
	}
	return tms, nil
}

const sqlStrToAddPost = "insert into post (userID,themeID,state,title,content,createdTime,lastEditTime,editTimes,praiseTimes,belittleTimes) values (?,?,?,?,?,?,?,?,?,?)"

//AddPost 新增帖子
func (s *sqlBase) AddPost(post *forum.Post) error {
	back, err := s.db.Exec(sqlStrToAddPost, post.UserID, post.ThemeID, post.State, post.Title, post.Content, post.CreatedTime, post.LastEditTime, post.EditTimes, post.PraiseTimes, post.BelittleTimes)
	if err != nil {
		return err
	}
	post.ID, err = back.LastInsertId()
	return err
}

const sqlStrToDeletePost = `
delete from comment where postID = ?;
delete from post where ID = ?`

//DeletePost 删除帖子，同时删除所有评论
func (s *sqlBase) DeletePost(postID int64) error {
	_, err := s.db.Exec(sqlStrToDeletePost, postID, postID)
	return err
}

const sqlStrToUpdatePost = "update post set state=?,title=?,content=?,lastEditTime=?,editTimes=?,praiseTimes=?,belittleTimes=? where ID=?"

//UpdatePost 更新帖子，只支持修改标题、内容、状态、最终编辑时间、编辑次数、赞、踩
func (s *sqlBase) UpdatePost(post *forum.Post) error {
	_, err := s.db.Exec(sqlStrToUpdatePost, post.State, post.Title, post.Content, post.LastEditTime, post.EditTimes, post.PraiseTimes, post.BelittleTimes, post.ID)
	return err
}

const sqlStrToQueryPostsSortType0 = `
select 
    p.ID,
    p.title,
    count(cmt.ID) as cmtCount,
    u1.ID,
    u1.name,
    p.createdTime,
    u2.ID,
    u2.name,
    max(cmt.createdTime)
from
    post as p,
    user as u1,
    comment as cmt,
    user as u2
where 
    p.themeID == ?
    and p.userID = u1.ID
    and cmt.postID = p.ID
    and u2.ID = cmt.userID
group by 
    p.ID
order by 
    p.createdTime desc
limit ? 
offset ?`

const sqlStrToQueryPostsSortType1 = `
select 
    p.ID,
    p.title,
    count(cmt.ID) as cmtCount,
    u1.ID,
    u1.name,
    p.createdTime,
    u2.ID,
    u2.name,
    max(cmt.createdTime)
from
    post as p,
    user as u1,
    comment as cmt,
    user as u2
where 
    p.themeID == ?
    and p.userID = u1.ID
    and cmt.postID = p.ID
    and u2.ID = cmt.userID
group by 
	p.ID
order by 
	max(cmt.createdTime) desc
limit ? 
offset ?`

//QueryPosts 查询帖子列表
func (s *sqlBase) QueryPosts(themeID int64, count, offset, sortType int) ([]*forum.PostBriefInfo, error) {
	var rows *sql.Rows
	var err error
	if sortType == 0 {
		rows, err = s.db.Query(sqlStrToQueryPostsSortType0, themeID, count, offset)
	} else {
		rows, err = s.db.Query(sqlStrToQueryPostsSortType1, themeID, count, offset)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := make([]*forum.PostBriefInfo, 0, count)
	for rows.Next() {
		post := new(forum.PostBriefInfo)
		err = rows.Scan(&post.ID, &post.Title, &post.CommentCount, &post.CreaterID, &post.CreaterName, &post.CreateTime, &post.LastCmterID, &post.LastCmterName, &post.LastCmtTime)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

const sqlStrToQueryPost = "select * from post where ID = ?"

//QueryPost 查询帖子 带所有评论的引用
func (s *sqlBase) QueryPost(postID int64) (*forum.Post, error) {
	post := new(forum.Post)
	err := s.db.QueryRow(sqlStrToQueryPost, postID).Scan(&post.ID, &post.UserID, &post.ThemeID, &post.State, &post.Title, &post.Content, &post.CreatedTime, &post.LastEditTime, &post.EditTimes, &post.PraiseTimes, &post.BelittleTimes)
	if err != nil {
		return nil, err
	}
	post.Comments, err = s.QueryComments(postID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

const sqlStrToAddComment = "insert into comment (userID,state,postID,content,createdTime,lastEditTime,editTimes,praiseTimes,belittleTimes) values (?,?,?,?,?,?,?,?,?)"

//AddComment 增加评论
func (s *sqlBase) AddComment(cmt *forum.Comment) error {
	back, err := s.db.Exec(sqlStrToAddComment, cmt.UserID, cmt.State, cmt.PostID, cmt.Content, cmt.CreatedTime, cmt.LastEditTime, cmt.EditTimes, cmt.PraiseTimes, cmt.BelittleTimes)
	if err != nil {
		return err
	}
	cmt.ID, err = back.LastInsertId()
	return err
}

const sqlStrToDeleteComment = "delete from comment where ID =?"

//DeleteComment 删除评论
func (s *sqlBase) DeleteComment(cmtID int64) error {
	_, err := s.db.Exec(sqlStrToDeleteComment, cmtID)
	return err
}

const sqlStrToUpdateComment = "update comment set state=?,content=?,lastEditTime=?,editTimes=?,praiseTimes=?,belittleTimes=? where ID=?"

//UpdateComment 更新评论，只支持修改内容、状态、最终编辑时间、编辑次数、赞、踩
func (s *sqlBase) UpdateComment(cmt *forum.Comment) error {
	_, err := s.db.Exec(sqlStrToUpdateComment, cmt.State, cmt.Content, cmt.LastEditTime, cmt.EditTimes, cmt.PraiseTimes, cmt.BelittleTimes, cmt.ID)
	return err
}

const sqlStrToQueryComments = "select * from comment where postID = ? order by createdTime"

//QueryComments 查询评论，按照创建时间排序
func (s *sqlBase) QueryComments(postID int64) ([]*forum.Comment, error) {
	rows, err := s.db.Query(sqlStrToQueryComments, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cmts := make([]*forum.Comment, 0, 20)
	for rows.Next() {
		cmt := new(forum.Comment)
		err = rows.Scan(&cmt.ID, &cmt.UserID, &cmt.State, &cmt.PostID, &cmt.Content, &cmt.CreatedTime, &cmt.LastEditTime, &cmt.EditTimes, &cmt.PraiseTimes, &cmt.BelittleTimes)
		if err != nil {
			return nil, err
		}
		cmts = append(cmts, cmt)
	}
	return cmts, nil
}

const sqlStrToAddUser = "insert into user (name,account,password,signUpTime,userType,state) values (?,?,?,?,?,?)"

//AddUser 新增用户
func (s *sqlBase) AddUser(user *forum.User) error {
	back, err := s.db.Exec(sqlStrToAddUser, user.Name, user.Account, user.PassWord, user.SignUpTime, user.UserType, user.UserState)
	if err != nil {
		return err
	}
	user.ID, err = back.LastInsertId()
	return err
}

const sqlStrToDeleteUser = "delete from user where ID =?"

//DeleteUser 删除用户
func (s *sqlBase) DeleteUser(userID int64) error {
	_, err := s.db.Exec(sqlStrToDeleteUser, userID)
	return err
}

const sqlStrToUpdateUser = "update user set name=?,password=?,userType=?,state=? where ID=?"

//UpdateUser 更新用户信息，只支持修改昵称、密码、用户类型、状态
func (s *sqlBase) UpdateUser(user *forum.User) error {
	_, err := s.db.Exec(sqlStrToUpdateUser, user.Name, user.PassWord, user.UserType, user.UserState, user.ID)
	return err
}

const sqlStrToQueryUserByID = "select * from user where ID = ?"

//QueryUserByID 查询用户
func (s *sqlBase) QueryUserByID(userID int64) (*forum.User, error) {
	user := new(forum.User)
	err := s.db.QueryRow(sqlStrToQueryUserByID, userID).Scan(&user.ID, &user.Name, &user.Account, &user.PassWord, &user.SignUpTime, &user.UserType, &user.UserState)
	if err != nil {
		err = &tool.ErrQueryNoResult{QueryItem: "用户"}
		return nil, err
	}
	return user, nil
}

const sqlStrToQueryUserByAccountAndPwd = "select * from user where account = ? and password = ?"

//QueryUserByAccountAndPwd 查询用户
func (s *sqlBase) QueryUserByAccountAndPwd(account string, password string) (*forum.User, error) {
	user := new(forum.User)
	err := s.db.QueryRow(sqlStrToQueryUserByAccountAndPwd, account, password).Scan(&user.ID, &user.Name, &user.Account, &user.PassWord, &user.SignUpTime, &user.UserType, &user.UserState)
	if err != nil {
		err = &tool.ErrQueryNoResult{QueryItem: "用户"}
		return nil, err
	}
	return user, nil
}

const sqlStrToIsUserNameExist = "select ID from user where name = ?"

//IsUserNameExist 是否昵称已存在
func (s *sqlBase) IsUserNameExist(name string) bool {
	row := s.db.QueryRow(sqlStrToIsUserNameExist, name)
	err := row.Scan(new(int64))
	return err == nil
}

const sqlStrToIsUserAccountExist = "select ID from user where account = ?"

//IsUserAccountExist 是否账号已存在
func (s *sqlBase) IsUserAccountExist(account string) bool {
	row := s.db.QueryRow(sqlStrToIsUserAccountExist, account)
	err := row.Scan(new(int64))
	return err == nil
}
