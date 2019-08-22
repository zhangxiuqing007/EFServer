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

//AddPost AddPost
func (s *SqliteIns) AddPost(post *forum.Post) error {
	sqlStr := "insert into post (userID,state,title,content,createdTime,lastEditTime,editTimes,praiseTimes,belittleTimes) values (?,?,?,?,?,?,?,?,?)"
	_, err := s.db.Exec(sqlStr, post.UserID, post.State, post.Title, post.Content, post.CreatedTime, post.LastEditTime, post.EditTimes, post.PraiseTimes, post.BelittleTimes)
	return err
}

//DeletePost DeletePost
func (s *SqliteIns) DeletePost(post *forum.Post) error {
	return nil
}

//UpdatePost UpdatePost
func (s *SqliteIns) UpdatePost(post *forum.Post) error {
	return nil
}

//QueryPost QueryPost
func (s *SqliteIns) QueryPost(id int64) (*forum.Post, error) {
	return nil, nil
}

//AddComment AddComment
func (s *SqliteIns) AddComment(cmt *forum.Comment) error {
	sqlStr := "insert into comment (userID,state,postID,content,createdTime,lastEditTime,editTimes,praiseTimes,belittleTimes) values (?,?,?,?,?,?,?,?,?)"
	_, err := s.db.Exec(sqlStr, cmt.UserID, cmt.State, cmt.PostID, cmt.Content, cmt.CreatedTime, cmt.LastEditTime, cmt.EditTimes, cmt.PraiseTimes, cmt.BelittleTimes)
	return err
}

//DeleteComment DeleteComment
func (s *SqliteIns) DeleteComment(comment *forum.Comment) error {
	return nil
}

//UpdateComment UpdateComment
func (s *SqliteIns) UpdateComment(comment *forum.Comment) error {
	return nil
}

//QueryComment QueryComment
func (s *SqliteIns) QueryComment(id int64) (*forum.Comment, error) {
	return nil, nil
}

//AddUser AddUser
func (s *SqliteIns) AddUser(user *forum.User) error {
	sqlStr := "insert into user (name,account,password,signUpTime,userType,state) values (?,?,?,?,?,?)"
	_, err := s.db.Exec(sqlStr, user.Name, user.Account, user.PassWord, user.SignUpTime, user.UserType, user.UserState)
	return err
}

//DeleteUser DeleteUser
func (s *SqliteIns) DeleteUser(user *forum.User) error {
	return nil
}

//UpdateUser UpdateUser
func (s *SqliteIns) UpdateUser(user *forum.User) error {
	return nil
}

//QueryUserByAccountAndPwd QueryUserByAccountAndPwd
func (s *SqliteIns) QueryUserByAccountAndPwd(account string, password string) (*forum.User, error) {
	row := s.db.QueryRow("select * from user where account = ? and password = ?", account, password)
	user := new(forum.User)
	err := row.Scan(&user.ID, &user.Name, &user.Account, &user.PassWord, &user.SignUpTime, &user.UserType, &user.UserState)
	if err != nil {
		return nil, tool.ErrQueryNoResult{QueryItem: "用户"}
	}
	return user, nil
}

//IsUserNameExist IsUserNameExist
func (s *SqliteIns) IsUserNameExist(name string) bool {
	row := s.db.QueryRow("select ID from user where name = ?", name)
	err := row.Scan(new(int64))
	return err == nil
}

//IsUserAccountExist IsUserAccountExist
func (s *SqliteIns) IsUserAccountExist(account string) bool {
	row := s.db.QueryRow("select ID from user where account = ?", account)
	err := row.Scan(new(int64))
	return err == nil
}
