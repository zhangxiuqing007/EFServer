package dba

import "EFServer/forum"

//DataOper DataOper must have value from main function
var DataOper IDataOper = &ioTool{}

//IDataOper DataSaver
type IDataOper interface {
	AddPost(post *forum.Post) error
	DeletePost(post *forum.Post) error
	UpdatePost(post *forum.Post) error

	AddComment(comment *forum.Comment) error
	DeleteComment(comment *forum.Comment) error
	UpdateComment(comment *forum.Comment) error

	AddUser(user *forum.User) error
	DeleteUser(user *forum.User) error
	UpdateUser(user *forum.User) error
}
