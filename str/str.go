package str

import (
	"fmt"
	"strconv"
	"strings"
	"math"
	"bytes"
)

// 字符串替换
func Replace(origin, search, replace string) string {
	return strings.Replace(origin, search, replace, -1)
}

// 使用map进行字符串替换
func ReplaceByMap(origin string, replaces map[string]string) string {
	result := origin
	for k, v := range replaces {
		result = strings.Replace(result, k, v, -1)
	}
	return result
}

// 字符串转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 字符串转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 字符串首字母转换为大写
func UcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	if IsLetterLower(s[0]) {
		return string(s[0]-32) + s[1:]
	}
	return s
}

// 字符串首字母转换为小写
func LcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	if IsLetterUpper(s[0]) {
		return string(s[0]+32) + s[1:]
	}
	return s
}

// 便利数组查找字符串索引位置，如果不存在则返回-1，使用完整遍历查找
func SearchArray(a []string, s string) int {
	for i, v := range a {
		if s == v {
			return i
		}
	}
	return -1
}

// 判断字符串是否在数组中
func InArray(a []string, s string) bool {
	return SearchArray(a, s) != -1
}

// 判断给定字符是否小写
func IsLetterLower(b byte) bool {
	if b >= byte('a') && b <= byte('z') {
		return true
	}
	return false
}

// 判断给定字符是否大写
func IsLetterUpper(b byte) bool {
	if b >= byte('A') && b <= byte('Z') {
		return true
	}
	return false
}

// 判断锁给字符串是否为数字
func IsNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < byte('0') || s[i] > byte('9') {
			return false
		}
	}
	return true
}

// 字符串截取，支持中文
func SubStr(str string, start int, length ...int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)
	// 简单的越界判断
	if start < 0 {
		start = 0
	}
	if start >= lth {
		start = lth
	}
	end := lth
	if len(length) > 0 {
		end = start + length[0]
		if end < start {
			end = lth
		}
	}
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[start:end])
}

// 字符串长度截取限制，超过长度限制被截取并在字符串末尾追加指定的内容，支持中文
func StrLimit(str string, length int, suffix ...string) (string) {
	rs := []rune(str)
	if len(str) < length {
		return str
	}
	addstr := "..."
	if len(suffix) > 0 {
		addstr = suffix[0]
	}
	return string(rs[0:length]) + addstr
}

// 按照百分比从字符串中间向两边隐藏字符(主要用于姓名、手机号、邮箱地址、身份证号等的隐藏)，支持utf-8中文，支持email格式。
func HideStr(str string, percent int, hide string) string {
	array := strings.Split(str, "@")
	if len(array) > 1 {
		str = array[0]
	}
	rs := []rune(str)
	length := len(rs)
	mid := math.Floor(float64(length / 2))
	hideLen := int(math.Floor(float64(length) * (float64(percent) / 100)))
	start := int(mid - math.Floor(float64(hideLen)/2))
	hideStr := []rune("")
	hideRune := []rune(hide)
	for i := 0; i < int(hideLen); i++ {
		hideStr = append(hideStr, hideRune...)
	}
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString(string(rs[0:start]))
	buffer.WriteString(string(hideStr))
	buffer.WriteString(string(rs[start+hideLen:]))
	if len(array) > 1 {
		buffer.WriteString(array[1])
	}
	return buffer.String()
}

// 将\n\r替换为html中的<br>标签。
func Nl2Br(str string) string {
	str = Replace(str, "\r\n", "\n")
	str = Replace(str, "\n\r", "\n")
	str = Replace(str, "\n", "<br />")
	return str
}

// 判断slice中是否包含某个字符串
func Contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// 转换数字ID到字符串
func ConvertID(intId int64) string {
	const mapping = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	code := ""
	size := int64(len(mapping))
	for intId >= size {
		mod := intId % size
		intId = intId / size

		code += mapping[mod : mod+1]
	}
	code += mapping[intId : intId+1]
	code = Reverse(code)

	return code
}

// 翻转字符串
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 对比版本号，返回-1，0，1三个值
func VersionCompare(version1 string, version2 string) int8 {
	if len(version1) == 0 {
		if len(version2) == 0 {
			return 0
		}

		return -1
	}

	if len(version2) == 0 {
		return 1
	}

	pieces1 := strings.Split(version1, ".")
	pieces2 := strings.Split(version2, ".")
	count1 := len(pieces1)
	count2 := len(pieces2)

	for i := 0; i < count1; i ++ {
		if i > count2-1 {
			return 1
		}

		piece1 := pieces1[i]
		piece2 := pieces2[i]
		len1 := len(piece1)
		len2 := len(piece2)

		if len1 == 0 {
			if len2 == 0 {
				continue
			}
		}

		maxLength := 0
		if len1 > len2 {
			maxLength = len1
		} else {
			maxLength = len2
		}

		piece1 = fmt.Sprintf("%0"+strconv.Itoa(maxLength)+"s", piece1)
		piece2 = fmt.Sprintf("%0"+strconv.Itoa(maxLength)+"s", piece2)

		if piece1 > piece2 {
			return 1
		}

		if piece1 < piece2 {
			return -1
		}
	}

	if count1 > count2 {
		return 1
	}

	if count1 == count2 {
		return 0
	}

	return -1
}
