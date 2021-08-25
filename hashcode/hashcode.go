package hashcode

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

// HashCode 自定义获取HashCode
func HashCode (did string) int {
	var didByte = []byte(did)
	byte2Md5 := fmt.Sprintf("%X", md5.Sum(didByte))
	var md52Byte = []byte(byte2Md5)
	hash2Int64 := int64(binary.LittleEndian.Uint64(md52Byte))
	return int(Abs(hash2Int64) % 100)
}

// Abs 绝对值
func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}