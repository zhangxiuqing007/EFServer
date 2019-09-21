package dba

import (
	"EFServer/forum"
	"database/sql"
	"errors"
)

type sqlBase struct {
	db *sql.DB
}

const sqlStrToClear = `
delete from comment;
delete from post;
delete from theme;
delete from user;
vacuum;`

func (s *sqlBase) Clear() error {
	_, err := s.db.Exec(sqlStrToClear)
	return err
}

//Close 关闭
func (s *sqlBase) Close() error {
	return s.db.Close()
}

const sqlStrToAddTheme = "insert into theme (name) values (?)"

//AddTheme 增加主题
func (s *sqlBase) AddTheme(themeName string) (*forum.ThemeInDB, error) {
	tmIns := new(forum.ThemeInDB)
	tmIns.Name = themeName
	result, err := s.db.Exec(sqlStrToAddTheme, themeName)
	if err != nil {
		return tmIns, err
	}
	tmIns.ID, err = result.LastInsertId()
	if err != nil {
		return tmIns, err
	}
	return tmIns, err
}

const sqlStrToDeleteTheme = "delete from theme where ID = ?"

//DeleteTheme 删除主题
func (s *sqlBase) DeleteTheme(themeID int64) error {
	_, err := s.db.Exec(sqlStrToDeleteTheme, themeID)
	return err
}

const sqlStrToUpdateTheme = "update theme set name = ? where ID = ?"

//UpdateTheme 更新主题
func (s *sqlBase) UpdateTheme(theme *forum.ThemeInDB) error {
	_, err := s.db.Exec(sqlStrToUpdateTheme, theme.Name, theme.ID)
	return err
}

const sqlStrToQueryTheme = "select name from theme where ID = ?"

//QueryTheme 查询某个主题
func (s *sqlBase) QueryTheme(themeID int64) (*forum.ThemeInDB, error) {
	tm := new(forum.ThemeInDB)
	tm.ID = themeID
	err := s.db.QueryRow(sqlStrToQueryTheme, themeID).Scan(&tm.Name)
	return tm, err
}

const sqlStrToQueryThemes = "select * from theme"

//QueryThemes 查询主题列表
func (s *sqlBase) QueryThemes() ([]*forum.ThemeInDB, error) {
	rows, err := s.db.Query(sqlStrToQueryThemes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tms := make([]*forum.ThemeInDB, 0, 20)
	for rows.Next() {
		tm := new(forum.ThemeInDB)
		err = rows.Scan(&tm.ID, &tm.Name)
		if err != nil {
			return tms, err
		}
		tms = append(tms, tm)
	}
	return tms, err
}

const sqlStrToAddPost = `insert into post (themeID,userID,title,state) values (?,?,?,?)`

//AddPost 新增帖子
func (s *sqlBase) AddPost(post *forum.PostInDB) error {
	back, err := s.db.Exec(sqlStrToAddPost,
		post.ThemeID,
		post.UserID,
		post.Title,
		post.State)
	if err != nil {
		return err
	}
	post.ID, err = back.LastInsertId()
	return err
}

//AddPosts 批量新增帖子
func (s *sqlBase) AddPosts(posts []*forum.PostInDB) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(sqlStrToAddPost)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, post := range posts {
		result, err := stmt.Exec(
			post.ThemeID,
			post.UserID,
			post.Title,
			post.State)
		if err != nil {
			return err
		}
		post.ID, err = result.LastInsertId()
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

const sqlStrToDeletePost = `
delete from comment where postID = ?;
delete from post where ID = ?`

//DeletePost 删除帖子，同时删除所有评论
func (s *sqlBase) DeletePost(postID int64) error {
	_, err := s.db.Exec(sqlStrToDeletePost, postID, postID)
	return err
}

const sqlStrToQueryPostCountOfTheme = "select count(post.ID) from post where post.themeID = ?"

//QueryPostCountOfTheme 统计本主题所有帖子总量
func (s *sqlBase) QueryPostCountOfTheme(themeID int64) int {
	var count int
	s.db.QueryRow(sqlStrToQueryPostCountOfTheme, themeID).Scan(&count)
	return count
}

const sqlStrToQueryPostsSortType0 = `
select 
    p.ID,
    p.title,
    count(cmt.ID) as cmtCount,
    u1.ID,
    u1.name,
    min(cmt.createdTime),
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
order by p.ID
limit ? 
offset ?`

const sqlStrToQueryPostsSortType1 = `
select 
    p.ID,
    p.title,
    count(cmt.ID) as cmtCount,
    u1.ID,
    u1.name,
    min(cmt.createdTime),
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
order max(cmt.createdTime) desc
limit ? 
offset ?`

//QueryPosts 查询某主题下的帖子列表
func (s *sqlBase) QueryPosts(themeID int64, count, offset, sortType int) ([]*forum.PostOnThemePage, error) {
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
	posts := make([]*forum.PostOnThemePage, 0, count)
	for rows.Next() {
		post := new(forum.PostOnThemePage)
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.CmtCount,
			&post.CreaterID,
			&post.CreaterName,
			&post.CreateTime,
			&post.LastCmterID,
			&post.LastCmterName,
			&post.LastCmtTime)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, err
}

const sqlStrToQueryPost = "select themeID,userID,title,state from post where ID = ?"

//QueryPost 查询帖子内容
func (s *sqlBase) QueryPost(postID int64) (*forum.PostInDB, error) {
	post := new(forum.PostInDB)
	post.ID = postID
	err := s.db.QueryRow(sqlStrToQueryPost, postID).Scan(
		&post.ThemeID,
		&post.UserID,
		&post.Title,
		&post.State)
	if err != nil {
		return nil, err
	}
	return post, nil
}

const sqlStrToQueryPostPG = `
select 
   p.title,
   tm.ID,
   tm.name
from 
    post as p
	join theme as tm
where 
	p.ID = ? and tm.ID = p.themeID`

//QueryPostPG 查询帖子内容
func (s *sqlBase) QueryPostPG(postID int64) (*forum.PostOnPostPage, error) {
	post := new(forum.PostOnPostPage)
	post.ID = postID
	err := s.db.QueryRow(sqlStrToQueryPostPG, postID).Scan(
		&post.Title,
		&post.ThemeID,
		&post.ThemeName)
	if err != nil {
		return nil, err
	}
	return post, nil
}

const sqlStrToAddComment = `
insert into comment 
(
	postID,
	userID,
	content,
	state,
	createdTime,
	lastEditTime,
	editTimes,
	praiseTimes,
	belittleTimes
)
values (?,?,?,?,?,?,?,?,?)`

//AddComment 增加评论
func (s *sqlBase) AddComment(cmt *forum.CommentInDB) error {
	back, err := s.db.Exec(sqlStrToAddComment,
		cmt.PostID,
		cmt.UserID,
		cmt.Content,
		cmt.State,
		cmt.CreatedTime,
		cmt.LastEditTime,
		cmt.EditTimes,
		cmt.PraiseTimes,
		cmt.BelittleTimes)
	if err != nil {
		return err
	}
	cmt.ID, err = back.LastInsertId()
	return err
}

//AddComments 批量增加评论
func (s *sqlBase) AddComments(comments []*forum.CommentInDB) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(sqlStrToAddComment)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, cmt := range comments {
		_, err := stmt.Exec(
			cmt.PostID,
			cmt.UserID,
			cmt.Content,
			cmt.State,
			cmt.CreatedTime,
			cmt.LastEditTime,
			cmt.EditTimes,
			cmt.PraiseTimes,
			cmt.BelittleTimes)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

const sqlStrToDeleteComment = "delete from comment where ID =?"

//DeleteComment 删除单个评论
func (s *sqlBase) DeleteComment(cmtID int64) error {
	_, err := s.db.Exec(sqlStrToDeleteComment, cmtID)
	return err
}

const sqlStrToQueryComments = "select * from comment where postID = ? order by createdTime"

//QueryComments 查询评论，按照创建时间排序
func (s *sqlBase) QueryComments(postID int64) ([]*forum.CommentInDB, error) {
	rows, err := s.db.Query(sqlStrToQueryComments, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cmts := make([]*forum.CommentInDB, 0, 50)
	for rows.Next() {
		cmt := new(forum.CommentInDB)
		err = rows.Scan(
			&cmt.ID,
			&cmt.PostID,
			&cmt.UserID,
			&cmt.Content,
			&cmt.State,
			&cmt.CreatedTime,
			&cmt.LastEditTime,
			&cmt.EditTimes,
			&cmt.PraiseTimes,
			&cmt.BelittleTimes)
		if err != nil {
			return nil, err
		}
		cmts = append(cmts, cmt)
	}
	return cmts, nil
}

const sqlStrToQueryPgComments = `
select 
       cmt.ID,
       cmt.content,
       cmt.praiseTimes,
       cmt.belittleTimes,
       u.ID,
       u.name,
       cmt.createdTime
from 
     comment as cmt
     join user as u
where 
	  cmt.postID = ? and cmt.userID = u.ID
order by cmt.ID`

//QueryPgComments 查询评论内容，用于显示在帖子页中
func (s *sqlBase) QueryPgComments(postID int64) ([]*forum.CmtOnPostPage, error) {
	cmts := make([]*forum.CmtOnPostPage, 0, 50)
	rows, err := s.db.Query(sqlStrToQueryPgComments, postID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cmt := new(forum.CmtOnPostPage)
		err = rows.Scan(
			&cmt.ID,
			&cmt.Content,
			&cmt.PraiseTimes,
			&cmt.BelittleTimes,
			&cmt.CmterID,
			&cmt.CmterName,
			&cmt.CmtTime)
		if err != nil {
			return nil, err
		}
		cmts = append(cmts, cmt)
	}
	return cmts, nil
}

const sqlStrToAddUser = `
insert into user 
(
	account,
	password,
	name,
	userType,
	state,
	signUpTime
)
values (?,?,?,?,?,?)`

//AddUser 新增用户
func (s *sqlBase) AddUser(user *forum.UserInDB) error {
	back, err := s.db.Exec(sqlStrToAddUser,
		user.Account,
		user.PassWord,
		user.Name,
		user.UserType,
		user.UserState,
		user.SignUpTime)
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

const sqlStrToQueryUserSaInfoByIDSection1 = `
select
      u.ID,
      u.name,
      u.signUpTime,
      u.userType,
      u.state,
      count(p.ID)
from
    user as u,
    post as p
where u.ID = ?
and p.userID = u.ID`

const sqlStrToQueryUserSaInfoByIDSection2 = `
select
      count(cmt.ID),
      sum(cmt.praiseTimes),
      sum(cmt.belittleTimes),
      max(cmt.lastEditTime)
from
    user as u,
    comment as cmt
where u.ID = ?
and cmt.userID = u.ID`

//QueryUserSaInfoByID 查询用户统计信息
func (s *sqlBase) QueryUserSaInfoByID(userID int64) (*forum.UserStatisticsInfo, error) {
	user := new(forum.UserStatisticsInfo)
	err := s.db.QueryRow(sqlStrToQueryUserSaInfoByIDSection1, userID).Scan(
		&user.ID,
		&user.Name,
		&user.SignUpTime,
		&user.UserType,
		&user.UserState,
		&user.PostTotalCount)
	if err != nil {
		return nil, err
	}
	err = s.db.QueryRow(sqlStrToQueryUserSaInfoByIDSection2, userID).Scan(
		&user.CmtTotalCount,
		&user.TotalPraisedTimes,
		&user.TotalBelittledTimes,
		&user.LastOperateTime)
	if err != nil {
		return nil, err
	}
	//总评论数减去发帖数，就是评论数
	user.CmtTotalCount -= user.PostTotalCount
	return user, err
}

const sqlStrToQueryUserByID = "select account,password,name,userType,state,signUpTime from user where ID = ?"

//QueryUserByID 查询用户
func (s *sqlBase) QueryUserByID(userID int64) (*forum.UserInDB, error) {
	user := new(forum.UserInDB)
	user.ID = userID
	err := s.db.QueryRow(sqlStrToQueryUserByID, userID).Scan(
		&user.Account,
		&user.PassWord,
		&user.Name,
		&user.UserType,
		&user.UserState,
		&user.SignUpTime)
	if err != nil {
		err = errors.New("无此用户")
		return nil, err
	}
	return user, nil
}

const sqlStrToQueryUserByAccountAndPwd = "select ID,name,userType,state,signUpTime from user where account = ? and password = ?"

//QueryUserByAccountAndPwd 查询用户
func (s *sqlBase) QueryUserByAccountAndPwd(account string, password string) (*forum.UserInDB, error) {
	user := new(forum.UserInDB)
	user.Account = account
	user.PassWord = password
	err := s.db.QueryRow(sqlStrToQueryUserByAccountAndPwd, account, password).Scan(
		&user.ID,
		&user.Name,
		&user.UserType,
		&user.UserState,
		&user.SignUpTime)
	if err != nil {
		err = errors.New("无此用户")
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
