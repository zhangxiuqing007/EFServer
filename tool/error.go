package tool

//StrError 基本的文字型错误
type StrError struct {
	ErrorStr string
}

func (err StrError) Error() string {
	return err.ErrorStr
}

//QueryNoResultError 查询无结果错误
type QueryNoResultError struct {
	QueryItem string
}

func (err QueryNoResultError) Error() string {
	return err.QueryItem + "查询失败"
}

//DataRepeatError 数据重复错误
type DataRepeatError struct {
	RepeatItem string
}

func (err DataRepeatError) Error() string {
	return err.RepeatItem + "重复"
}
