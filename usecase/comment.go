package usecase

import (
	"EFServer/dba"
	"EFServer/forum"
	"time"
)

//CommentingData io data to add comment
//build by controller
type CommentingData struct {
	UserID  uint64
	PostID  uint64
	Content string
}

func (data *CommentingData) buildCommentIns() *forum.Comment {
	comment := new(forum.Comment)
	comment.ID = 0
	comment.UserID = data.UserID
	comment.PostID = data.PostID
	comment.Content = data.Content
	comment.CreatedTime = time.Now().UnixNano()
	comment.LastEditTime = comment.CreatedTime
	comment.EditTimes = 0
	comment.PraiseTimes = 0
	comment.BelittleTimes = 0
	return comment
}

//AddComment add comment
//call by controller
func AddComment(data *CommentingData) error {
	comment := data.buildCommentIns()
	return dba.DataOper.AddComment(comment)
}
