package usecase

import "EFServer/forum"

//GetAllThemes 获取所有的主题指针
func GetAllThemes() ([]*forum.ThemeInDB, error) {
	//先从数据库读取
	return db.QueryThemes()
}

//GetTheme 获取主题
func GetTheme(themeID int64) (*forum.ThemeInDB, error) {
	return db.QueryTheme(themeID)
}
