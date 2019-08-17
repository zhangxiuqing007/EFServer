package usecase

import "testing"

func TestAddNewUser(t *testing.T) {
	newUser := new(UserSignUpData)
	newUser.Name = "我最大最强"
	newUser.Account = "690313521@qq.com"
	newUser.Password = "aa135828"
	err := AddUser(newUser)
	if err != nil {
		t.Error(err)
	}
}
