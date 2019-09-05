package usecase

import "EFServer/forum"

//GetAllThemes 获取所有的主题指针
func GetAllThemes() ([]*forum.Theme, error) {
	//先从数据库读取
	return db.QueryThemes()
}

//GetTheme 获取主题
func GetTheme(themeName string) (*forum.Theme, error) {
	return db.QueryTheme(themeName)
}
