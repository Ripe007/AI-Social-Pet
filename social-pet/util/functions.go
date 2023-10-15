package util

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
)

func Md5String(InputStr string) string {
	return Md5([]byte(InputStr))
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}
