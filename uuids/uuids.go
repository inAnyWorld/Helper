package uuids

import (
	uuid "github.com/nu7hatch/gouuid"
	"log"
)

// Uuid 生成uuid
func Uuid() string {
	fileId, uuidErr := uuid.NewV4()
	if uuidErr != nil {
		log.Println("uuid 生成Err:", uuidErr)
		return "false"
	}
	return fileId.String()
}
