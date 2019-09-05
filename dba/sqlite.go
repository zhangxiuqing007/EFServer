package dba

import (
	"EFServer/forum"
	"EFServer/tool"
	"database/sql"

	//sqlite3 driver
	_ "github.com/mattn/sqlite3"
)

//SqliteIns sqlite实现
type SqliteIns struct {
	db *sql.DB
}

//Open 打开
func (s *SqliteIns) Open(dbFilePath string) error {
	var err error
	s.db, err = sql.Open("sqlite3", dbFilePath)
	return err
}

//Close 关闭
func (s *SqliteIns) Close() error {
	return s.db.Close()
}

//AddTheme 增加主题
func (s *SqliteIns) AddTheme(themeName string) error {
	sqlStr := "insert into theme (name) values (?)"
	_, err := s.db.Exec(sqlStr, themeName)
	return err
}

//DeleteTheme 删除主题
func (s *SqliteIns) DeleteTheme(themeID int64) error {
	sqlStr := "delete from theme where ID = ?"
	_, err := s.db.Exec(sqlStr, themeID)
	return err
}

//UpdateTheme 更新主题
func (s *SqliteIns) UpdateTheme(theme *forum.Theme) error {
	sqlStr := "update theme set name = ? where ID = ?"
	_, err := s.db.Exec(sqlStr, theme.Name, theme.ID)
	return err
}

//QueryTheme 查询某个主题
func (s *SqliteIns) QueryTheme(themeName string) (*forum.Theme, error) {
	sqlStr := "select * from theme where name = ?"
	row := s.db.QueryRow(sqlStr, themeName)
	tm := new(forum.Theme)
	err := row.Scan(&tm.ID, &tm.Name)
	if err != nil {
		return nil, err
	}
	return tm, nil
}

//QueryThemes 查询主题列表
func (s *SqliteIns) QueryThemes() ([]*forum.Theme, error) {
	sqlStr := "select * from theme"
	rows, err := s.db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	back := make([]*forum.Theme, 0, 10)
	for rows.Next() {
		tm := forum.Theme{}
		err = rows.Scan(&tm.ID, &tm.Name)
		if err != nil {
			return nil, err
		}
		back = append(back, &tm)
	}
	return back, nil
}

//AddPost 新增帖子
func (s *SqliteIns) AddPost(post *forum.Post) error {
	sqlStr := "insert into post (userID,themeID,state,title,content,createdTime,lastEditTime,editTimes,praiseTimes,belittleTimes) values (?,?,?,?,?,?,?,?,?,?)"
	back, err := s.db.Exec(sqlStr, post.UserID, post.ThemeID, post.State, post.Title, post.Content, post.CreatedTime, post.LastEditTime, post.EditTimes, post.PraiseTimes, post.BelittleTimes)
	if err != nil {
		return err
	}
	post.ID, err = back.LastInsertId()
	return err
}

//DeletePost 删除帖子，同时删除所有评论
func (s *SqliteIns) DeletePost(postID int64) error {
	sqlStr := "delete from comment where postID = ?; delete from post where ID = ?"
	_, err := s.db.Exec(sqlStr, postID, postID)
	return err
}

//UpdatePost 更新帖子，只支持修改标题、内容、状态、最终编辑时间、编辑次数、赞、踩
func (s *SqliteIns) UpdatePost(post *forum.Post) error {
	sqlStr := "update post set state=?,title=?,content=?,lastEditTime=?,editTimes=?,praiseTimes=?,belittleTimes=? where ID=?"
	_, err := s.db.Exec(sqlStr, post.State, post.Title, post.Content, post.LastEditTime, post.EditTimes, post.PraiseTimes, post.BelittleTimes, post.ID)
	return err
}

//QueryPosts 查询帖子列表
func (s *SqliteIns) QueryPosts(themeID int64) ([]*forum.PostBriefInfo, error) {
	posts := make([]*forum.PostBriefInfo, 0, 100)
	sqlStr := "select * from post p left join user u on p.userID = u.ID where p.themeID = ?"
	rows, err := s.db.Query(sqlStr, themeID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post := new(forum.PostBriefInfo)
		if err = rows.Scan( /* 未完成 */ ); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

//QueryPost 查询帖子 带所有评论的引用
func (s *SqliteIns) QueryPost(postID int64) (*forum.Post, error) {
	post := new(forum.Post)
	sqlStr := "select * from post where ID = ?"
	err := s.db.QueryRow(sqlStr, postID).Scan(&post.ID, &post.UserID, &post.ThemeID, &post.State, &post.Title, &post.Content, &post.CreatedTime, &post.LastEditTime, &post.EditTimes, &post.PraiseTimes, &post.BelittleTimes)
	if err != nil {
		return nil, err
	}
	post.Comments, err = s.QueryComments(postID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

//AddComment 增加评论
func (s *SqliteIns) AddComment(cmt *forum.Comment) error {
	sqlStr := "insert into comment (userID,state,postID,content,createdTime,lastEditTime,editTimes,praiseTimes,belittleTimes) values (?,?,?,?,?,?,?,?,?)"
	back, err := s.db.Exec(sqlStr, cmt.UserID, cmt.State, cmt.PostID, cmt.Content, cmt.CreatedTime, cmt.LastEditTime, cmt.EditTimes, cmt.PraiseTimes, cmt.BelittleTimes)
	if err != nil {
		return err
	}
	cmt.ID, err = back.LastInsertId()
	return err
}

//DeleteComment 删除评论
func (s *SqliteIns) DeleteComment(cmtID int64) error {
	sqlStr := "delete from comment where ID =?"
	_, err := s.db.Exec(sqlStr, cmtID)
	return err
}

//UpdateComment 更新评论，只支持修改内容、状态、最终编辑时间、编辑次数、赞、踩
func (s *SqliteIns) UpdateComment(cmt *forum.Comment) error {
	sqlStr := "update comment set state=?,content=?,lastEditTime=?,editTimes=?,praiseTimes=?,belittleTimes=? where ID=?"
	_, err := s.db.Exec(sqlStr, cmt.State, cmt.Content, cmt.LastEditTime, cmt.EditTimes, cmt.PraiseTimes, cmt.BelittleTimes, cmt.ID)
	return err
}

//QueryComments 查询评论，按照创建时间排序
func (s *SqliteIns) QueryComments(postID int64) ([]*forum.Comment, error) {
	sqlStr := "select * from comment where postID = ? order by createdTime"
	rows, err := s.db.Query(sqlStr, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cmts := make([]*forum.Comment, 0, 10)
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

//AddUser 新增用户
func (s *SqliteIns) AddUser(user *forum.User) error {
	sqlStr := "insert into user (name,account,password,signUpTime,userType,state) values (?,?,?,?,?,?)"
	back, err := s.db.Exec(sqlStr, user.Name, user.Account, user.PassWord, user.SignUpTime, user.UserType, user.UserState)
	if err != nil {
		return err
	}
	user.ID, err = back.LastInsertId()
	return err
}

//DeleteUser 删除用户
func (s *SqliteIns) DeleteUser(userID int64) error {
	sqlStr := "delete from user where ID =?"
	_, err := s.db.Exec(sqlStr, userID)
	return err
}

//UpdateUser 更新用户信息，只支持修改昵称、密码、用户类型、状态
func (s *SqliteIns) UpdateUser(user *forum.User) error {
	sqlStr := "update user set name=?,password=?,userType=?,state=? where ID=?"
	_, err := s.db.Exec(sqlStr, user.Name, user.PassWord, user.UserType, user.UserState, user.ID)
	return err
}

//QueryUserByID 查询用户
func (s *SqliteIns) QueryUserByID(userID int64) (*forum.User, error) {
	row := s.db.QueryRow("select * from user where ID = ?", userID)
	user := new(forum.User)
	err := row.Scan(&user.ID, &user.Name, &user.Account, &user.PassWord, &user.SignUpTime, &user.UserType, &user.UserState)
	if err != nil {
		return nil, &tool.ErrQueryNoResult{QueryItem: "用户"}
	}
	return user, nil
}

//QueryUserByAccountAndPwd 查询用户
func (s *SqliteIns) QueryUserByAccountAndPwd(account string, password string) (*forum.User, error) {
	row := s.db.QueryRow("select * from user where account = ? and password = ?", account, password)
	user := new(forum.User)
	err := row.Scan(&user.ID, &user.Name, &user.Account, &user.PassWord, &user.SignUpTime, &user.UserType, &user.UserState)
	if err != nil {
		return nil, &tool.ErrQueryNoResult{QueryItem: "用户"}
	}
	return user, nil
}

//IsUserNameExist 是否昵称已存在
func (s *SqliteIns) IsUserNameExist(name string) bool {
	row := s.db.QueryRow("select ID from user where name = ?", name)
	err := row.Scan(new(int64))
	return err == nil
}

//IsUserAccountExist 是否账号已存在
func (s *SqliteIns) IsUserAccountExist(account string) bool {
	row := s.db.QueryRow("select ID from user where account = ?", account)
	err := row.Scan(new(int64))
	return err == nil
}
