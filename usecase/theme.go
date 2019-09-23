package usecase

import "EFServer/forum"

//QueryAllThemes 获取所有的主题指针
func QueryAllThemes() ([]*forum.ThemeInDB, error) {
	//先从数据库读取
	return db.QueryAllThemes()
}

//QueryTheme 获取主题
func QueryTheme(themeID int64) (*forum.ThemeInDB, error) {
	return db.QueryTheme(themeID)
}
