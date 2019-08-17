package tool

import "github.com/satori/uuid"

//NewUUID 创建一个全新的UUID
func NewUUID() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}
