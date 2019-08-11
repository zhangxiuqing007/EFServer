package usecase

import (
	"EFServer/dba"
	"EFServer/forum"
	"time"
)

//PostingData io data to add new post.
//this instance is from controller
type PostingData struct {
	UserID  uint64
	Title   string
	Content string
}

func (data *PostingData) buildPostIns() *forum.Post {
	post := new(forum.Post)
	post.ID = 0
	post.UserID = data.UserID
	post.Title = data.Title
	post.Content = data.Content
	post.CreatedTime = time.Now().UnixNano()
	post.LastEditTime = post.CreatedTime
	post.EditTimes = 0
	post.PraiseTimes = 0
	post.BelittleTimes = 0
	post.Comments = make([]forum.Comment, 0, 2)
	return post
}

//AddPost add a new post
//call by controller
func AddPost(data *PostingData) error {
	//check user authority

	//check post content and title legal

	//build post instance
	post := data.buildPostIns()
	//save
	return dba.DataOper.AddPost(post)
}

//QueryPost query a post content
//call by controller
func QueryPost(id uint64) (post *forum.Post, err error) {
	post, err = dba.DataOper.QueryPost(id)
	return
}
