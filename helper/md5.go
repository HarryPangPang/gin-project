package helper

import (
	"crypto/md5"
	"fmt"
)

func Md5Decode(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
