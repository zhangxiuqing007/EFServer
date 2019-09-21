package usecase

import (
	"EFServer/forum"
)

//QueryPost 帖子查询
func QueryPost(postID int64) (*forum.PostInDB, error) {
	return db.QueryPost(postID)
}

//QueryPosts 查询帖子列表
func QueryPosts(themeID int64, count, offset, sortType int) ([]*forum.PostOnThemePage, error) {
	return db.QueryPosts(themeID, count, offset, sortType)
}

//QueryPostPG 帖子页内容查询
func QueryPostPG(postID int64) (*forum.PostOnPostPage, error) {
	return db.QueryPostPG(postID)
}

//QueryPostCountOfTheme 获取该主题内帖子的总数量
func QueryPostCountOfTheme(themeID int64) int {
	return db.QueryPostCountOfTheme(themeID)
}
