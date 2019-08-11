package usecase

import "testing"
import "time"
import "github.com/satori/uuid"

func TestAddNewPost(t *testing.T) {
	for i := 0; i < 2; i++ {
		newPost := new(PostingData)
		newPost.UserID = 4
		guidTitle, _ := uuid.NewV4()
		newPost.Title = guidTitle.String() + ": " + getNowTimeStr()
		guidContent, _ := uuid.NewV4()
		newPost.Content = guidContent.String() + ": " + getNowTimeStr()
		err := AddPost(newPost)
		if err != nil {
			t.Error(err)
		}
	}
}

func getNowTimeStr() string {
	return time.Now().String()
}
