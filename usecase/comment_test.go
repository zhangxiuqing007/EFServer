package usecase

import "testing"
import "github.com/satori/uuid"

func TestAddNewCmt(t *testing.T) {
	for i := 0; i < 2; i++ {
		newCmt := new(CommentingData)
		newCmt.UserID = 4
		newCmt.PostID = 4
		guidContent, _ := uuid.NewV4()
		newCmt.Content = guidContent.String() + ": " + getNowTimeStr()
		err := AddComment(newCmt)
		if err != nil {
			t.Error(err)
		}
	}
}
