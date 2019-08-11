package forum

//Post PostClass
type Post struct {
	PostBase
	Title    string
	Comments []Comment
}

const (
	//PostStateNormal NormalState
	PostStateNormal = iota
	//PostStateLock LockState
	PostStateLock
)
