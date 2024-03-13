package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// InList 判断key是否存在于列表中
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// Md5 hash检验图片内容是否重复
func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
