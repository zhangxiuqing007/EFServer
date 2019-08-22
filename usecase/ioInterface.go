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

	AddPost(post *forum.Post) error
	DeletePost(post *forum.Post) error
	UpdatePost(post *forum.Post) error
	QueryPost(id int64) (*forum.Post, error)

	AddComment(comment *forum.Comment) error
	DeleteComment(comment *forum.Comment) error
	UpdateComment(comment *forum.Comment) error
	QueryComment(id int64) (*forum.Comment, error)

	AddUser(user *forum.User) error
	DeleteUser(user *forum.User) error
	UpdateUser(user *forum.User) error
	QueryUserByAccountAndPwd(account string, password string) (*forum.User, error)

	IsUserNameExist(name string) bool
	IsUserAccountExist(account string) bool
}
