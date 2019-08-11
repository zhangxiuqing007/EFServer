package dba

import (
	"EFServer/forum"
	"EFServer/tool"
	"database/sql"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func linkToDb() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/efdb?charset=utf8")
	return
}

type ioTool struct {
}

func (t ioTool) AddPost(post *forum.Post) error {
	db, err := linkToDb()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStr := "insert into post (userID,state,title,content,createdTime,lastEditTime,editTimes,praiseTimes,belittleTimes) values (?,?,?,?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, post.UserID, post.State, post.Title, post.Content, post.CreatedTime, post.LastEditTime, post.EditTimes, post.PraiseTimes, post.BelittleTimes)
	return err
}

func (t ioTool) DeletePost(post *forum.Post) error {
	return &tool.StrError{ErrorStr: "DeletePost...Not Completed Now!"}
}

func (t ioTool) UpdatePost(post *forum.Post) error {
	return &tool.StrError{ErrorStr: "UpdatePost...Not Completed Now!"}
}

func (t ioTool) QueryPost(id uint64) (*forum.Post, error) {
	return nil, &tool.StrError{ErrorStr: "UpdatePost...Not Completed Now!"}
}

func (t ioTool) AddComment(cmt *forum.Comment) error {
	db, err := linkToDb()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStr := "insert into comment (userID,state,postID,content,createdTime,lastEditTime,editTimes,praiseTimes,belittleTimes) values (?,?,?,?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, cmt.UserID, cmt.State, cmt.PostID, cmt.Content, cmt.CreatedTime, cmt.LastEditTime, cmt.EditTimes, cmt.PraiseTimes, cmt.BelittleTimes)
	return err
}

func (t ioTool) DeleteComment(cmt *forum.Comment) error {
	return &tool.StrError{ErrorStr: "DeleteComment...Not Completed Now!"}
}

func (t ioTool) UpdateComment(cmt *forum.Comment) error {
	return &tool.StrError{ErrorStr: "UpdateComment...Not Completed Now!"}
}

func (t ioTool) AddUser(user *forum.User) error {
	db, err := linkToDb()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStr := "insert into user (name,account,password,signUpTime,userType,state) values (?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, user.Name, user.Account, user.PassWord, user.SignUpTime, user.UserType, user.UserState)
	return err
}

func (t ioTool) DeleteUser(user *forum.User) error {
	return &tool.StrError{ErrorStr: "DeleteUser...Not Completed Now!"}
}

func (t ioTool) UpdateUser(user *forum.User) error {
	return &tool.StrError{ErrorStr: "UpdateUser...Not Completed Now!"}
}

func (t ioTool) QueryUserByCodeAndPwd(account string, password string) (*forum.User, error) {
	return nil, &tool.StrError{ErrorStr: "QueryUserByCodeAndPwd...Not Completed Now!"}
}

func (t ioTool) IsUserNameExist(name string) (bool, error) {
	return t.isUserFieldExsit("name", name)
}

func (t ioTool) IsAccountExist(account string) (bool, error) {
	return t.isUserFieldExsit("account", account)
}

func (t ioTool) isUserFieldExsit(feild string, patten string) (bool, error) {
	db, err := linkToDb()
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
