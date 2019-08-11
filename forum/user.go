package forum

//User UserInfo
type User struct {
	ID         uint64
	Name       string
	Account    string
	PassWord   string
	SignUpTime int64
	UserType   int
	UserState  int
}

const (
	//UserTypeAdministrator Administrator
	UserTypeAdministrator = iota
	//UserTypeNormalUser NormalUser
	UserTypeNormalUser
)

const (
	//UserStateNormal normal
	UserStateNormal = iota
	//UserStateLock lock
	UserStateLock
)
