package usecase

import (
	"EFServer/forum"
	"errors"
	"time"
	"unicode/utf8"
)

//UserSignUpData 新用户注册传输用数据结构，由Controller创建。
type UserSignUpData struct {
	Name     string
	Account  string
	Password string
}

func (data UserSignUpData) buildUserIns() *forum.UserInDB {
	user := new(forum.UserInDB)
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
	if utf8.RuneCountInString(data.Name) == 0 {
		return errors.New("昵称不合法（至少一个字）")
	}
	//检查账户合法性
	if utf8.RuneCountInString(data.Account) < 3 {
		return errors.New("账号不合法（至少三个字符）")
	}
	//检查密码合法性
	if utf8.RuneCountInString(data.Password) < 3 {
		return errors.New("密码不合法（至少三个字符）")
	}
	//检查昵称占用
	if db.IsUserNameExist(data.Name) {
		return errors.New("昵称被占用")
	}
	//检查账户占用
	if db.IsUserAccountExist(data.Account) {
		return errors.New("账号被占用")
	}
	//保存
	user := data.buildUserIns()
	return db.AddUser(user)
}

//QueryUser 用户查询
func QueryUser(account string, password string) (*forum.UserInDB, error) {
	return db.QueryUserByAccountAndPwd(account, password)
}

//QueryUserSaInfo 查询用户统计信息
func QueryUserSaInfo(userID int64) (*forum.UserStatisticsInfo, error) {
	return db.QueryUserSaInfoByID(userID)
}
