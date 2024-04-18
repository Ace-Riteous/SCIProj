package utils

import (
	"SCIProj/global"
	"crypto/md5"
	"fmt"
	"strings"
)

//给字符串生成md5
//@params str 需要加密的字符串
//@params salt interface{} 加密的盐
//@return str 返回md5码

func Md5Crypt(str string) (CryptStr string) {
	if l := len(global.MD5SALT); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), global.MD5SALT)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
