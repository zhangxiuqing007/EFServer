package forum

//User UserInfo
type User struct {
	ID       uint64
	UserCode string
	PassWord string
	UserType int
}

const (
	//Administrator Administrator
	Administrator = iota
	//NormalUser NormalUser
	NormalUser
)
