package usecase

import (
	"EFServer/forum"
)

//QueryPost 帖子查询
func QueryPost(postID int64) (*forum.PostInDB, error) {
	return db.QueryPost(postID)
}

//QueryPostsOfTheme 查询帖子列表
func QueryPostsOfTheme(themeID int64, count, offset, sortType int) ([]*forum.PostOnThemePage, error) {
	return db.QueryPostsOfTheme(themeID, count, offset, sortType)
}

//QueryPostsOfUser 查询某个用户发的帖子的列表
func QueryPostsOfUser(userID int64, count, offset int) ([]*forum.PostOnThemePage, error) {
	return db.QueryPostsOfUser(userID, count, offset)
}

//QueryPostOfPostPage 帖子页内容查询
func QueryPostOfPostPage(postID int64) (*forum.PostOnPostPage, error) {
	return db.QueryPostOfPostPage(postID)
}

//QueryPostCountOfTheme 获取该主题内帖子的总数量
func QueryPostCountOfTheme(themeID int64) (int, error) {
	return db.QueryPostCountOfTheme(themeID)
}

//QueryPostCountOfUser 获取该用户的发帖总量
func QueryPostCountOfUser(userID int64) (int, error) {
	return db.QueryPostCountOfUser(userID)
}
