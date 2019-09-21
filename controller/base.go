package controller

type loginInfo struct {
	IsLogin  bool
	UserName string
}

//获取提供的导航页
func getNaviPageIndexs(
	currentPageIndex int, /*当前页索引*/
	countOnePage int, /*一页元素数量*/
	maxHalfNaviPageCount int, /*最大的导航页数量的一半*/
	elementTotalCount int) /*元素的总数量*/ (beginIndex, endIndex int) {
	//先计算beginIndex
	beginIndex = currentPageIndex - maxHalfNaviPageCount
	if beginIndex < 0 {
		beginIndex = 0
	}
	//再计算endIndex
	//剩余的元素数量
	leftElementCount := elementTotalCount - (currentPageIndex+1)*countOnePage
	//剩余的页数量
	leftPageCount := leftElementCount / countOnePage
	if leftElementCount%countOnePage > 0 {
		leftPageCount++
	}
	//剩余页上限
	if leftPageCount > maxHalfNaviPageCount {
		leftPageCount = maxHalfNaviPageCount
	}
	endIndex = currentPageIndex + leftPageCount
	return
}
