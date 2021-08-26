package gofile

import (
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// IsFileExist 文件是否存在
func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

// FileWrite 写文件
func FileWrite(fileName, content string) (err error) {
	f, err := os.Create(fileName)
	defer f.Close()
	if err == nil {
		_, err = f.WriteString(content)
	}
	return
}

// FileGet 读取文件
func FileGet(fileName string) (content string, err error) {
	f, err := ioutil.ReadFile(fileName)
	if err == nil {
		return string(f), nil
	}
	return "", err
}

// GetExt 获取文件扩展名
func GetExt(filePath string) string {
	suffix := filepath.Ext(filePath)
	if suffix != "" {
		return strings.ToLower(suffix[1:])
	}
	return suffix
}

// Basename 返回路径中的文件名部分
func Basename(filePath string) string {
	return filepath.Base(filePath)
}

// GetMime 获取mime类型
func GetMime(filePath string, fast bool) string {
	var res string
	if fast {
		suffix := filepath.Ext(filePath)
		res = mime.TypeByExtension(suffix)
	} else {
		srcFile, err := os.Open(filePath)
		if err != nil {
			return res
		}

		buffer := make([]byte, 512)
		_, err = srcFile.Read(buffer)
		if err != nil {
			return res
		}

		res = http.DetectContentType(buffer)
	}

	return res
}

// ReadInArray 文件写入数组
func ReadInArray(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

// ReadFile 读取文件内容
func ReadFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	return data, err
}

// FileSize 获取文件大小
func FileSize(filePath string) int64 {
	f, err := os.Stat(filePath)
	if nil != err {
		return -1
	}
	return f.Size()
}

// Rename 重命名
func Rename(oldName string, newName string) error {
	return os.Rename(oldName, newName)
}
