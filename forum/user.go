package forum

//UserStatisticsInfo 用户统计信息页展示内容
type UserStatisticsInfo struct {
	UserBase
	PostTotalCount      int
	CmtTotalCount       int
	TotalPraisedTimes   int
	TotalBelittledTimes int
	LastOperateTime     int64
}

//UserInDB 用户对象，在数据库中的表示
type UserInDB struct {
	UserBase
	Account  string
	PassWord string
}

//UserBase 用户基本信息
type UserBase struct {
	ID   int64
	Name string

	UserType  int
	UserState int

	SignUpTime int64
}

//GetUserTypeShowName 获取用户类型名称
func GetUserTypeShowName(userType int) string {
	switch userType {
	case UserTypeAdministrator:
		return "管理员"
	case UserTypeNormalUser:
		return "普通用户"
	default:
		return "错误类型"
	}
}

const (
	//UserTypeAdministrator 用户类型：管理员
	UserTypeAdministrator = iota
	//UserTypeNormalUser 用户类型：普通用户
	UserTypeNormalUser
)

//GetUserStateShowName 获取用户类型名称
func GetUserStateShowName(state int) string {
	switch state {
	case UserStateNormal:
		return "正常"
	case UserStateLock:
		return "锁定"
	default:
		return "错误状态"
	}
}

const (
	//UserStateNormal 用户状态：正常
	UserStateNormal = iota
	//UserStateLock 用户账号：锁定
	UserStateLock
)
