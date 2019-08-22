package usecase

import (
	"EFServer/forum"
	"time"
)

//PostingData 新增帖子传输用数据结构，由Controller创建。
type PostingData struct {
	UserID  int64
	Title   string
	Content string
}

func (data *PostingData) buildPostIns() *forum.Post {
	post := new(forum.Post)
	post.ID = 0
	post.UserID = data.UserID
	post.Title = data.Title
	post.Content = data.Content
	post.State = forum.PostStateNormal
	post.CreatedTime = time.Now().UnixNano()
	post.LastEditTime = post.CreatedTime
	post.EditTimes = 0
	post.PraiseTimes = 0
	post.BelittleTimes = 0
	post.Comments = make([]forum.Comment, 0, 0)
	return post
}

//AddPost 新增帖子
func AddPost(data *PostingData) error {
	//检查用户存在性、权限、状态
	//检查帖子标题合法性
	//检查帖子内容合法性
	//保存
	post := data.buildPostIns()
	return db.AddPost(post)
}

//QueryPost 帖子查询
func QueryPost(id int64) (*forum.Post, error) {
	return db.QueryPost(id)
}
