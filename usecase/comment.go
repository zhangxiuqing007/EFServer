package usecase

import (
	"EFServer/forum"
	"time"
)

//CommentingData 新增评论传输数据结构，只能由Controller来创建
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
	comment.State = forum.CmtStateNormal
	comment.CreatedTime = time.Now().UnixNano()
	comment.LastEditTime = comment.CreatedTime
	comment.EditTimes = 0
	comment.PraiseTimes = 0
	comment.BelittleTimes = 0
	return comment
}

//AddComment 新增评论
func AddComment(data *CommentingData) error {
	//检查用户存在性、状态、权限
	//检查评论内容合法
	//检查帖子存在性、状态
	//保存
	comment := data.buildCommentIns()
	return db.AddComment(comment)
}
