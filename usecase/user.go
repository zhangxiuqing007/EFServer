package usecase

import "EFServer/forum"
import "EFServer/tool"
import "time"

//UserSignUpData 新用户注册传输用数据结构，由Controller创建。
type UserSignUpData struct {
	Name     string
	Account  string
	Password string
}

func (data UserSignUpData) buildUserIns() *forum.User {
	user := new(forum.User)
	user.ID = 0
	user.Name = data.Name
	user.Account = data.Account
	user.PassWord = data.Password
	user.SignUpTime = time.Now().UnixNano()
	user.UserType = forum.UserTypeNormalUser
	user.UserState = forum.UserStateNormal
	return user
}

//AddUser signUp
func AddUser(data *UserSignUpData) error {
	//检查昵称合法性
	//检查账户合法性
	//检查密码合法性
	//检查昵称占用
	exist, err := db.IsUserNameExist(data.Name)
	if err != nil {
		return err
	}
	if exist {
		return tool.DataRepeatError{RepeatItem: "昵称"}
	}
	//检查账户占用
	exist, err = db.IsAccountExist(data.Account)
	if err != nil {
		return err
	}
	if exist {
		return tool.DataRepeatError{RepeatItem: "账号"}
	}
	//保存
	user := data.buildUserIns()
	return db.AddUser(user)
}

//QueryUser 用户查询
func QueryUser(account string, password string) (*forum.User, error) {
	return db.QueryUserByAccountAndPwd(account, password)
}
