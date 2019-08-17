package dba

import (
	"EFServer/forum"
	"EFServer/tool"
	"database/sql"

	//sqlite3 driver
	_ "github.com/mattn/sqlite3"
)

//SqliteDbFilePath SqliteDbFilePath
var SqliteDbFilePath string

func linkToSqlite() (*sql.DB, error) {
	return sql.Open("sqlite3", SqliteDbFilePath)
}

//SqliteIns sqlite实现
type SqliteIns struct {
}

//AddPost AddPost
func (s *SqliteIns) AddPost(post *forum.Post) error {
	db, err := linkToSqlite()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStr := "insert into post (userID,state,title,content,createdTime,lastEditTime,editTimes,praiseTimes,belittleTimes) values (?,?,?,?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, post.UserID, post.State, post.Title, post.Content, post.CreatedTime, post.LastEditTime, post.EditTimes, post.PraiseTimes, post.BelittleTimes)
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
func (s *SqliteIns) QueryPost(id uint64) (*forum.Post, error) {
	return nil, nil
}

//AddComment AddComment
func (s *SqliteIns) AddComment(cmt *forum.Comment) error {
	db, err := linkToSqlite()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStr := "insert into comment (userID,state,postID,content,createdTime,lastEditTime,editTimes,praiseTimes,belittleTimes) values (?,?,?,?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, cmt.UserID, cmt.State, cmt.PostID, cmt.Content, cmt.CreatedTime, cmt.LastEditTime, cmt.EditTimes, cmt.PraiseTimes, cmt.BelittleTimes)
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

//AddUser AddUser
func (s *SqliteIns) AddUser(user *forum.User) error {
	db, err := linkToSqlite()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStr := "insert into user (name,account,password,signUpTime,userType,state) values (?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, user.Name, user.Account, user.PassWord, user.SignUpTime, user.UserType, user.UserState)
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
	db, err := linkToSqlite()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("select * from user where account = ? and password = ?", account, password)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		user := new(forum.User)
		err = rows.Scan(&user.ID, &user.Name, &user.Account, &user.PassWord, &user.SignUpTime, &user.UserType, &user.UserState)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, tool.QueryNoResultError{QueryItem: "账号"}
}

//IsUserNameExist IsUserNameExist
func (s *SqliteIns) IsUserNameExist(name string) (bool, error) {
	return s.isUserFieldExsit("name", name)
}

//IsAccountExist IsAccountExist
func (s *SqliteIns) IsAccountExist(account string) (bool, error) {
	return s.isUserFieldExsit("account", account)
}

//isUserFieldExsit isUserFieldExsit
func (s *SqliteIns) isUserFieldExsit(feild string, patten string) (bool, error) {
	db, err := linkToSqlite()
	if err != nil {
		return false, err
	}
	defer db.Close()
	rows, err := db.Query("select ID from user where ? = ?", feild, patten)
	defer rows.Close()
	if err != nil {
		return false, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}
