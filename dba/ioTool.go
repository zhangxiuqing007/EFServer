package dba

import (
	"EFServer/forum"
	"fmt"
)

//IoTool use to full interface
type ioTool struct {
}

//AddPost  AddPost
func (tool ioTool) AddPost(post *forum.Post) error {
	fmt.Println("AddPost...Not Completed Now!")
	return nil
}

//DeletePost DeletePost
func (tool ioTool) DeletePost(post *forum.Post) error {
	fmt.Println("DeletePost...Not Completed Now!")
	return nil
}

//UpdatePost UpdatePost
func (tool ioTool) UpdatePost(post *forum.Post) error {
	fmt.Println("UpdatePost...Not Completed Now!")
	return nil
}

//AddComment  AddComment
func (tool ioTool) AddComment(comment *forum.Comment) error {
	fmt.Println("AddComment...Not Completed Now!")
	return nil
}

//DeleteComment DeleteComment
func (tool ioTool) DeleteComment(comment *forum.Comment) error {
	fmt.Println("DeleteComment...Not Completed Now!")
	return nil
}

//UpdateComment UpdateComment
func (tool ioTool) UpdateComment(comment *forum.Comment) error {
	fmt.Println("UpdateComment...Not Completed Now!")
	return nil
}

//AddUser AddUser
func (tool ioTool) AddUser(user *forum.User) error {
	fmt.Println("AddUser...Not Completed Now!")
	return nil
}

//DeleteUser DeleteUser
func (tool ioTool) DeleteUser(user *forum.User) error {
	fmt.Println("DeleteUser...Not Completed Now!")
	return nil
}

//UpdateUser UpdateUser
func (tool ioTool) UpdateUser(user *forum.User) error {
	fmt.Println("UpdateUser...Not Completed Now!")
	return nil
}
