package tool

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unicode/utf8"
)

func TestUtf8StringCount(t *testing.T) {
	str := "中文"
	if utf8.RuneCountInString(str) != 2 {
		t.Error("错误的工作函数：utf8.RuneCountInString")
		t.FailNow()
	}
	str2 := "11中文字符串"
	if utf8.RuneCountInString(str2) != 7 {
		t.Errorf("错误的工作函数：utf8.RuneCountInString")
		t.FailNow()
	}
}

func Test_RandomValue(t *testing.T) {
	ints := make([]int, 0, 1000)
	i := 0
	for {
		ints = append(ints, rand.Intn(100))
		i++
		if i > 1000 {
			break
		}
	}
}

func Test_FormatTime(t *testing.T) {
	ticks := time.Now().UnixNano()
	timeIns := time.Unix(0, ticks)
	if ticks != timeIns.UnixNano() {
		t.Error("时间不一致")
		t.FailNow()
	} else {
		t.Log("时间一致")
	}
	t.Log(timeIns.String())
	t.Log(timeIns.Format("yyyy-mm-dd hh-mm-ss"))
	t.Log(fmt.Sprintf("%d年%d月%d日 %s %d:%d:%d",
		timeIns.Year(),
		timeIns.Month(),
		timeIns.Day(),
		weekdayStrs[timeIns.Weekday()],
		timeIns.Hour(),
		timeIns.Minute(),
		timeIns.Second()))
}
