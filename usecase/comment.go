package usecase

import (
	"EFServer/forum"
)

//QueryCommentsOfPostPage 查询评论内容，用户帖子页内展示
func QueryCommentsOfPostPage(postID int64, count, offset int) ([]*forum.CmtOnPostPage, error) {
	return db.QueryCommentsOfPostPage(postID, count, offset)
}

//QueryCommentsCountOfPost 查询帖子的评论总量
func QueryCommentsCountOfPost(postID int64) (int, error) {
	return db.QueryCommentsCountOfPost(postID)
}
