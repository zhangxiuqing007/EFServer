package tool

//StrError base error define
type StrError struct {
	ErrorStr string
}

func (err StrError) Error() string {
	return err.ErrorStr
}

//QueryNoResultError query from db but has no result
type QueryNoResultError struct {
	QueryContent string
}

func (err QueryNoResultError) Error() string {
	return err.QueryContent + " query......has no result"
}

//DataRepeatError has repeat data
type DataRepeatError struct {
	RepeatContent string
}

func (err DataRepeatError) Error() string {
	return err.RepeatContent + " data is repeat"
}
