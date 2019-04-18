package rand

import (
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var digits  = []rune("0123456789")

// 自定义的 rand.Intn
func intn (max int) int {
	return int(time.Now().UnixNano())%max
}

// 获得一个 min, max 之间的随机数(min <= x <= max)
func Rand (min, max int) int {
	if min >= max {
		return min
	}
	if min == 0 {
		return intn(max + 1)
	}
	if min > 0 {
		// 数值往左平移，再使用底层随机方法获得随机数，随后将结果数值往右平移
		return intn(max - (min - 0) + 1) + (min - 0)
	}
	if min < 0 {
		// 数值往右平移，再使用底层随机方法获得随机数，随后将结果数值往左平移
		return intn(max + (0 - min) + 1) - (0 - min)
	}
	return 0
}

// 获得指定长度的随机字符串(可能包含数字和字母)
func RandStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		if intn(2) == 1 {
			b[i] = digits[intn(10)]
		} else {
			b[i] = letters[intn(52)]
		}
	}
	return string(b)
}

// 获得指定长度的随机数字字符串
func RandDigits(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = digits[intn(10)]
	}
	return string(b)
}

// 获得指定长度的随机字母字符串
func RandLetters(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[intn(52)]
	}
	return string(b)
}