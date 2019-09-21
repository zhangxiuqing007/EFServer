package tool

import (
	"errors"
	"math/rand"
	"time"
)

//SplitText 拆分文本
func SplitText(s string, spe []rune) []string {
	if len(s) == 0 || spe == nil || len(spe) == 0 {
		panic(errors.New("参数错误"))
	}
	runes := append([]rune(s), spe[0])
	result := make([]string, 0, len(runes)/4)
	//是否包含方法
	isContain := func(r rune) bool {
		for _, v := range spe {
			if v == r {
				return true
			}
		}
		return false
	}
	index := -1
	for i, v := range runes {
		if isContain(v) {
			if index >= 0 {
				result = append(result, string(runes[index:i]))
				index = -1
			}
		} else {
			if index == -1 {
				index = i
			}
		}
	}
	return result
}

var firstNameWords []string
var lastNameWords []string

//InitNameWords 初始化随机名字用字符
func InitNameWords(f, l []string) {
	firstNameWords = f
	lastNameWords = l
}

func getRandomFirstNameWord() []rune {
	return []rune(firstNameWords[rand.Intn(len(firstNameWords))])
}
func getRandomLastNameWord() []rune {
	return []rune(lastNameWords[rand.Intn(len(lastNameWords))])
}

//RandomChineseName 随机生成中文名字
func RandomChineseName() string {
	rand.Seed(time.Now().UnixNano())
	name := make([]rune, 0, 5)
	name = append(name, getRandomFirstNameWord()...)
	name = append(name, getRandomLastNameWord()...)
	//70%的人名字是两个字
	if rand.Intn(100) > 30 {
		name = append(name, getRandomLastNameWord()...)
	}
	return string(name)
}
