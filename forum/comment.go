package forum

//Comment Comment Class
type Comment struct {
	PostBase
	PostID uint64
}

const (
	//CmtStateNormal normal
	CmtStateNormal = iota
	//CmtStateLock lock
	CmtStateLock
)
