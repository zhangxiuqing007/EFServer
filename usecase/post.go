package usecase

import (
	"EFServer/forum"
	"time"
)

//PostingData 新增帖子传输用数据结构，由Controller创建。
type PostingData struct {
	ThemeID int64
	UserID  int64
	Title   string
	Content string
}

func (data *PostingData) buildPostIns() *forum.Post {
	post := new(forum.Post)
	post.ID = 0
	post.ThemeID = data.ThemeID
	post.UserID = data.UserID
	post.Title = data.Title
	post.Content = data.Content
	post.State = forum.PostStateNormal
	post.CreatedTime = time.Now().UnixNano()
	post.LastEditTime = post.CreatedTime
	post.EditTimes = 0
	post.PraiseTimes = 0
	post.BelittleTimes = 0
	//post.Comments = make([]*forum.Comment, 0, 0) //新帖子，无评论
	return post
}

//AddPost 新增帖子
func AddPost(data *PostingData) error {
	//检查主题存在性、状态
	//检查用户存在性、权限、状态
	//检查帖子标题合法性
	//检查帖子内容合法性
	//保存
	post := data.buildPostIns()
	return db.AddPost(post)
}

//QueryPost 帖子查询
func QueryPost(postID int64) (*forum.Post, error) {
	return db.QueryPost(postID)
}

//QueryPosts 查询帖子列表
func QueryPosts(themeID int64, count, offset, sortType int) ([]*forum.PostBriefInfo, error) {
	return db.QueryPosts(themeID, count, offset, sortType)
}
