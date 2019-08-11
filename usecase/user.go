package usecase

import "EFServer/forum"
import "EFServer/tool"
import "EFServer/dba"
import "time"

//UserSignUpData signup
type UserSignUpData struct {
	Name     string
	Account  string
	Password string
	UserType int
}

func (data UserSignUpData) buildUserIns() *forum.User {
	user := new(forum.User)
	user.ID = 0
	user.Name = data.Name
	user.Account = data.Account
	user.PassWord = data.Password
	user.UserType = data.UserType
	user.SignUpTime = time.Now().UnixNano()
	user.UserState = forum.UserStateNormal
	return user
}

//AddUser signUp
func AddUser(data *UserSignUpData) error {
	//check data legal, data length

	//check name repeat
	exist, err := dba.DataOper.IsUserNameExist(data.Name)
	if err != nil {
		return err
	}
	if exist {
		return tool.DataRepeatError{RepeatContent: "name"}
	}
	//check userCode repeat
	exist, err = dba.DataOper.IsAccountExist(data.Account)
	if err != nil {
		return err
	}
	if exist {
		return tool.DataRepeatError{RepeatContent: "account"}
	}
	//make instance
	user := data.buildUserIns()
	//save to db
	return dba.DataOper.AddUser(user)
}

//QueryUser QueryUser
func QueryUser(account string, password string) (*forum.User, error) {
	return dba.DataOper.QueryUserByCodeAndPwd(account, password)
}
