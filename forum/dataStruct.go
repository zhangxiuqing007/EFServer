package forum

//Post 帖子
type Post struct {
	PostBase
	Title    string
	Comments []Comment
}

//Comment 评论
type Comment struct {
	PostBase
	PostID uint64
}

// PostBase 基础结构
type PostBase struct {
	ID     uint64
	UserID uint64

	Content string

	State int

	CreatedTime  int64
	LastEditTime int64
	EditTimes    int

	PraiseTimes   int
	BelittleTimes int
}

//User 用户
type User struct {
	ID   uint64
	Name string

	Account  string
	PassWord string

	SignUpTime int64

	UserType int

	UserState int
}

const (
	//UserTypeAdministrator 用户类型：管理员
	UserTypeAdministrator = iota
	//UserTypeNormalUser 用户类型：普通用户
	UserTypeNormalUser
	//UserStateNormal 用户状态：正常
	UserStateNormal = iota
	//UserStateLock 用户账号：锁定
	UserStateLock
	//PostStateNormal 帖子状态：正常
	PostStateNormal = iota
	//PostStateLock 帖子状态：锁定
	PostStateLock
	//CmtStateNormal 评论状态：正常
	CmtStateNormal = iota
	//CmtStateLock 评论状态：锁定
	CmtStateLock
)
