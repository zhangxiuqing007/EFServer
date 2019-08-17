package usecase

import (
	"EFServer/tool"
	"testing"
)

func TestAddNewCmt(t *testing.T) {
	for i := 0; i < 5; i++ {
		newCmt := new(CommentingData)
		newCmt.UserID = 4
		newCmt.PostID = 4
		newCmt.Content = tool.NewUUID() + ": " + getNowTimeStr()
		err := AddComment(newCmt)
		if err != nil {
			t.Error(err)
		}
	}
}
