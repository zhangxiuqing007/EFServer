package forum

//Post PostClass
type Post struct {
	PostBase
	Comments []Comment
}

//Comment Comment Class
type Comment struct {
	PostBase
}

// PostBase ContentBaseClass
type PostBase struct {
	ID            uint64
	CreaterID     uint64
	Content       string
	CreateTime    int64
	LastEditTime  int64
	EditTimes     int
	PraiseTimes   int
	BelittleTimes int
}

const (
	//PostStateNormal NormalState
	PostStateNormal = iota
	//PostStateLock LockState
	PostStateLock
)
