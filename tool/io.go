package tool

import (
	"io/ioutil"
	"os"
)

//MustStr 一旦有错误就引起崩溃
func MustStr(str string, err error) string {
	if err != nil {
		panic(err)
	}
	return str
}

//ReadAllTextUtf8 读取utf8编码文本文件的全部内容并转换成string
func ReadAllTextUtf8(filePath string) (string, error) {
	buffer, err := ReadAllBytes(filePath)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

//MustBytes 一旦有错误就引起崩溃
func MustBytes(buffer []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return buffer
}

//ReadAllBytes 读取文件所有字节
func ReadAllBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
