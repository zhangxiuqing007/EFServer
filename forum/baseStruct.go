package forum

// PostBase ContentBaseClass
type PostBase struct {
	ID            uint64
	UserID        uint64
	Content       string
	State         int
	CreatedTime   int64
	LastEditTime  int64
	EditTimes     int
	PraiseTimes   int
	BelittleTimes int
}
