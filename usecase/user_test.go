package usecase

import "EFServer/forum"
import "testing"

func TestAddNewUser(t *testing.T) {
	newUser := new(UserSignUpData)
	newUser.Name = "我最大最强3"
	newUser.Account = "690313523@qq.com"
	newUser.Password = "aa135828"
	newUser.UserType = forum.UserTypeAdministrator
	err := AddUser(newUser)
	if err != nil {
		t.Error(err)
	}
}
