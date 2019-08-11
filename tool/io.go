package tool

import (
	"io/ioutil"
	"os"
)

//ReadFileString ReadFileString
func ReadFileString(path string) string {
	file, _ := os.Open(path)
	buffer, _ := ioutil.ReadAll(file)
	return string(buffer)
}
