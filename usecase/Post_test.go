package usecase

import "testing"
import "time"
import "EFServer/tool"

func TestAddNewPost(t *testing.T) {
	for i := 0; i < 2; i++ {
		newPost := new(PostingData)
		newPost.UserID = 0
		newPost.Title = tool.NewUUID() + ": " + getNowTimeStr()
		newPost.Content = tool.NewUUID() + ": " + getNowTimeStr()
		err := AddPost(newPost)
		if err != nil {
			t.Error(err)
		}
	}
}

func getNowTimeStr() string {
	return time.Now().String()
}
