package dba

import (
	"EFServer/forum"
	"EFServer/tool"
	"database/sql"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func linkToMysql() (*sql.DB, error) {
	return sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/efdb?charset=utf8")
}

//MySQLIns Mysql的IO实现
type MySQLIns struct {
}

//AddPost AddPost
func (m *MySQLIns) AddPost(post *forum.Post) error {
	db, err := linkToMysql()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStr := "insert into post (userID,state,title,content,createdTime,lastEditTime,editTimes,praiseTimes,belittleTimes) values (?,?,?,?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, post.UserID, post.State, post.Title, post.Content, post.CreatedTime, post.LastEditTime, post.EditTimes, post.PraiseTimes, post.BelittleTimes)
	return err
}

//DeletePost DeletePost
func (m *MySQLIns) DeletePost(post *forum.Post) error {
	return &tool.StrError{ErrorStr: "DeletePost...Not Completed Now!"}
}

//UpdatePost UpdatePost
func (m *MySQLIns) UpdatePost(post *forum.Post) error {
	return &tool.StrError{ErrorStr: "UpdatePost...Not Completed Now!"}
}

//QueryPost QueryPost
func (m *MySQLIns) QueryPost(id uint64) (*forum.Post, error) {
	return nil, &tool.StrError{ErrorStr: "UpdatePost...Not Completed Now!"}
}

//AddComment AddComment
func (m *MySQLIns) AddComment(cmt *forum.Comment) error {
	db, err := linkToMysql()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStr := "insert into comment (userID,state,postID,content,createdTime,lastEditTime,editTimes,praiseTimes,belittleTimes) values (?,?,?,?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, cmt.UserID, cmt.State, cmt.PostID, cmt.Content, cmt.CreatedTime, cmt.LastEditTime, cmt.EditTimes, cmt.PraiseTimes, cmt.BelittleTimes)
	return err
}

//DeleteComment DeleteComment
func (m *MySQLIns) DeleteComment(cmt *forum.Comment) error {
	return &tool.StrError{ErrorStr: "DeleteComment...Not Completed Now!"}
}

//UpdateComment UpdateComment
func (m *MySQLIns) UpdateComment(cmt *forum.Comment) error {
	return &tool.StrError{ErrorStr: "UpdateComment...Not Completed Now!"}
}

//AddUser AddUser
func (m *MySQLIns) AddUser(user *forum.User) error {
	db, err := linkToMysql()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStr := "insert into user (name,account,password,signUpTime,userType,state) values (?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, user.Name, user.Account, user.PassWord, user.SignUpTime, user.UserType, user.UserState)
	return err
}

//DeleteUser DeleteUser
func (m *MySQLIns) DeleteUser(user *forum.User) error {
	return &tool.StrError{ErrorStr: "DeleteUser...Not Completed Now!"}
}

//UpdateUser UpdateUser
func (m *MySQLIns) UpdateUser(user *forum.User) error {
	return &tool.StrError{ErrorStr: "UpdateUser...Not Completed Now!"}
}

//QueryUserByAccountAndPwd QueryUserByAccountAndPwd
func (m *MySQLIns) QueryUserByAccountAndPwd(account string, password string) (*forum.User, error) {
	return nil, &tool.StrError{ErrorStr: "QueryUserByCodeAndPwd...Not Completed Now!"}
}

//IsUserNameExist IsUserNameExist
func (m *MySQLIns) IsUserNameExist(name string) (bool, error) {
	return m.isUserFieldExsit("name", name)
}

//IsAccountExist IsAccountExist
func (m *MySQLIns) IsAccountExist(account string) (bool, error) {
	return m.isUserFieldExsit("account", account)
}

//isUserFieldExsit isUserFieldExsit
func (m *MySQLIns) isUserFieldExsit(feild string, patten string) (bool, error) {
	db, err := linkToMysql()
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
