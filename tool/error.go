package tool

//ErrStr 基本的文字型错误
type ErrStr struct {
	ErrorStr string
}

func (err *ErrStr) Error() string {
	return err.ErrorStr
}

//ErrQueryNoResult 查询无结果
type ErrQueryNoResult struct {
	QueryItem string
}

func (err *ErrQueryNoResult) Error() string {
	return err.QueryItem + "查询失败"
}

//ErrDataRepeat 数据重复错误
type ErrDataRepeat struct {
	RepeatItem string
}

func (err *ErrDataRepeat) Error() string {
	return err.RepeatItem + "重复"
}

//ErrUnDone 开发未完成
type ErrUnDone struct {
	UnDoneItem string
}

func (err *ErrUnDone) Error() string {
	return err.UnDoneItem + "：开发未完成"
}
