package forum

//Theme 主题
type Theme struct {
	ID   int64
	Name string
}

//PostBriefInfo 帖子简要内容
type PostBriefInfo struct {
	ID    int64
	Title string

	CreaterID   int64
	CreaterName string
	CreateTime  int64

	CommentCount  int
	LastCmterID   int64
	LastCmterName string
	LastCmtTime   int64

	PraiseTimes   int
	BelittleTimes int
}

//Post 帖子
type Post struct {
	PostBase
	Title    string
	ThemeID  int64
	Comments []*Comment
}

//Comment 评论
type Comment struct {
	PostBase
	PostID int64
}

// PostBase 基础结构
type PostBase struct {
	ID int64

	UserID int64

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
	ID   int64
	Name string

	Account  string
	PassWord string

	SignUpTime int64

	UserType  int
	UserState int
}

const (
	//UserTypeAdministrator 用户类型：管理员
	UserTypeAdministrator = iota
	//UserTypeNormalUser 用户类型：普通用户
	UserTypeNormalUser
)

const (
	//UserStateNormal 用户状态：正常
	UserStateNormal = iota
	//UserStateLock 用户账号：锁定
	UserStateLock
)

const (
	//PostStateNormal 帖子状态：正常
	PostStateNormal = iota
	//PostStateLock 帖子状态：锁定
	PostStateLock
	//PostStateHide 帖子状态：隐藏
	PostStateHide
)

const (
	//CmtStateNormal 评论状态：正常
	CmtStateNormal = iota
	//CmtStateLock 评论状态：锁定
	CmtStateLock
	//CmtStateHide 评论状态：隐藏
	CmtStateHide
)
