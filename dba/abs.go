package dba

import "EFServer/forum"

//Abs 测试
type Abs struct {
}

//AddPost 测试
func (a *Abs) AddPost(post *forum.Post) error {
	return nil
}

//DeletePost 测试
func (a *Abs) DeletePost(post *forum.Post) error {
	return nil
}

//UpdatePost 测试
func (a *Abs) UpdatePost(post *forum.Post) error {
	return nil
}

//QueryPost 测试
func (a *Abs) QueryPost(id uint64) (*forum.Post, error) {
	return nil, nil
}

//AddComment 测试
func (a *Abs) AddComment(comment *forum.Comment) error {
	return nil
}

//DeleteComment 测试
func (a *Abs) DeleteComment(comment *forum.Comment) error {
	return nil
}

//UpdateComment 测试
func (a *Abs) UpdateComment(comment *forum.Comment) error {
	return nil
}

//AddUser 测试
func (a *Abs) AddUser(user *forum.User) error {
	return nil
}

//DeleteUser 测试
func (a *Abs) DeleteUser(user *forum.User) error {
	return nil
}

//UpdateUser 测试
func (a *Abs) UpdateUser(user *forum.User) error {
	return nil
}

//QueryUserByAccountAndPwd 测试
func (a *Abs) QueryUserByAccountAndPwd(account string, password string) (*forum.User, error) {
	return &forum.User{
		Name:      "大力",
		Account:   "asd",
		PassWord:  "asd",
		UserType:  forum.UserTypeAdministrator,
		UserState: forum.UserStateNormal,
	}, nil
}

//IsUserNameExist 测试
func (a *Abs) IsUserNameExist(name string) (bool, error) {
	return false, nil
}

//IsAccountExist 测试
func (a *Abs) IsAccountExist(account string) (bool, error) {
	return false, nil
}
