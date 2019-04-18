package md5

import (
	"fmt"
	"crypto/md5"
	"os"
	"io"
)

// 将任意类型的变量进行md5摘要(注意map等非排序变量造成的不同结果)
func Encrypt(v string) string {
	h := md5.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 将字符串进行MD5哈希摘要计算
func EncryptString(v string) string {
	h := md5.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 将文件内容进行MD5哈希摘要计算
func EncryptFile(path string) string {
	f, e := os.Open(path)
	if e != nil {
		return ""
	}
	h := md5.New()
	_, e = io.Copy(h, f)
	if e != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
