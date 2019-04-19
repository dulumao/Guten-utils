// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// 时间管理
package time

import (
	"time"
	"regexp"
	"strings"
	"strconv"
	"errors"
	"fmt"
)

const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month

	TIME_REAGEX_PATTERN = `(\d{4}[-/]\d{2}[-/]\d{2})[\sT]{0,1}(\d{2}:\d{2}:\d{2}){0,1}\.{0,1}(\d{0,9})([\sZ]{0,1})([\+-]{0,1})([:\d]*)`
)

var (
	// 用于time.Time转换使用，防止多次Compile
	timeRegex *regexp.Regexp
)

func init() {
	// 使用正则判断会比直接使用ParseInLocation挨个轮训判断要快很多
	timeRegex, _ = regexp.Compile(TIME_REAGEX_PATTERN)

}

// 类似与js中的SetTimeout，一段时间后执行回调函数
func SetTimeout(t time.Duration, callback func()) {
	go func() {
		time.Sleep(t)
		callback()
	}()
}

// 类似与js中的SetInterval，每隔一段时间后执行回调函数，当回调函数返回true，那么继续执行，否则终止执行，该方法是异步的
// 注意：由于采用的是循环而不是递归操作，因此间隔时间将会以上一次回调函数执行完成的时间来计算
func SetInterval(t time.Duration, callback func() bool) {
	go func() {
		for {
			time.Sleep(t)
			if !callback() {
				break
			}
		}
	}()
}

// 设置当前进程全局的默认时区
func SetTimeZone(zone string) error {
	location, err := time.LoadLocation(zone)
	if err == nil {
		time.Local = location
	}
	return err
}

// 获取当前的纳秒数
func Nanosecond() int64 {
	return time.Now().UnixNano()
}

// 获取当前的微秒数
func Microsecond() int64 {
	return time.Now().UnixNano() / 1e3
}

// 获取当前的毫秒数
func Millisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

// 获取当前的秒数(时间戳)
func Second() int64 {
	return time.Now().UnixNano() / 1e9
}

// 获得当前的日期(例如：2006-01-02)
func Date() string {
	return time.Now().Format("2006-01-02")
}

// 获得当前的时间(例如：2006-01-02 15:04:05)
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 字符串转换为时间对象，支持的标准时间格式：
// "2017-12-14 04:51:34 +0805 LMT",
// "2006-01-02T15:04:05Z07:00",
// "2014-01-17T01:19:15+08:00",
// "2018-02-09T20:46:17.897Z",
// "2018-02-09 20:46:17.897",
// "2018-02-09T20:46:17Z",
// "2018-02-09 20:46:17",
// "2018-02-09",
func StrToTime(str string) (time.Time, error) {
	var result time.Time
	var local = time.Local
	if match := timeRegex.FindStringSubmatch(str); len(match) > 0 {
		var year, month, day, hour, min, sec, nsec int
		var array []string
		// 日期
		array = strings.Split(match[1], "-")
		if len(array) >= 3 {
			year, _ = strconv.Atoi(array[0])
			month, _ = strconv.Atoi(array[1])
			day, _ = strconv.Atoi(array[2])
		}
		// 时间
		array = strings.Split(match[2], ":")
		if len(array) >= 3 {
			hour, _ = strconv.Atoi(array[0])
			min, _ = strconv.Atoi(array[1])
			sec, _ = strconv.Atoi(array[2])
		}
		array = strings.Split(match[1], "-")
		// 纳秒，检查病执行位补齐
		if match[3] != "" {
			nsec, _ = strconv.Atoi(match[3])
			for i := 0; i < 9-len(match[3]); i++ {
				nsec *= 10
			}
		}
		// 如果字符串中有时区信息，那么执行时区转换，将时区转成UTC
		if match[4] != "" && match[6] == "" {
			match[6] = "000000"
		}
		if match[6] != "" {
			zone := strings.Replace(match[6], ":", "", -1)
			zone = strings.TrimLeft(zone, "+-")
			zone += strings.Repeat("0", 6-len(zone))
			h, _ := strconv.Atoi(zone[0:2])
			m, _ := strconv.Atoi(zone[2:4])
			s, _ := strconv.Atoi(zone[4:6])
			// 判断字符串输入的时区是否和当前程序时区相等，不相等则将对象统一转换为UTC时区
			// 当前程序时区Offset(秒)
			_, localOffset := time.Now().Zone()
			if (h*3600 + m*60 + s) != localOffset {
				local = time.UTC
				// UTC时差转换
				operation := match[5]
				if operation != "+" && operation != "-" {
					operation = "-"
				}
				switch operation {
				case "+":
					if h > 0 {
						hour -= h
					}
					if m > 0 {
						min -= m
					}
					if s > 0 {
						sec -= s
					}
				case "-":
					if h > 0 {
						hour += h
					}
					if m > 0 {
						min += m
					}
					if s > 0 {
						sec += s
					}
				}
			}
		}
		// 生成UTC时间对象
		result = time.Date(year, time.Month(month), day, hour, min, sec, nsec, local)
		return result, nil
	}
	return result, errors.New("unsupported time format")
}

// 字符串转换为时间对象，指定字符串时间格式，format格式形如：Y-m-d H:i:s
func StrToTimeFormat(str string, format string) (time.Time, error) {
	return StrToTimeLayout(str, formatToStdLayout(format))
}

// 字符串转换为时间对象，通过标准库layout格式进行解析，layout格式形如：2006-01-02 15:04:05
func StrToTimeLayout(str string, layout string) (time.Time, error) {
	if t, err := time.ParseInLocation(layout, str, time.Local); err == nil {
		return t, nil
	} else {
		return time.Time{}, err
	}
}

func computeTimeDiff(diff int64) (int64, string) {
	diffStr := ""
	switch {
	case diff <= 0:
		diff = 0
		diffStr = "now"
	case diff < 2:
		diff = 0
		diffStr = "1 second"
	case diff < 1*Minute:
		diffStr = fmt.Sprintf("%d seconds", diff)
		diff = 0

	case diff < 2*Minute:
		diff -= 1 * Minute
		diffStr = "1 minute"
	case diff < 1*Hour:
		diffStr = fmt.Sprintf("%d minutes", diff/Minute)
		diff -= diff / Minute * Minute

	case diff < 2*Hour:
		diff -= 1 * Hour
		diffStr = "1 hour"
	case diff < 1*Day:
		diffStr = fmt.Sprintf("%d hours", diff/Hour)
		diff -= diff / Hour * Hour

	case diff < 2*Day:
		diff -= 1 * Day
		diffStr = "1 day"
	case diff < 1*Week:
		diffStr = fmt.Sprintf("%d days", diff/Day)
		diff -= diff / Day * Day

	case diff < 2*Week:
		diff -= 1 * Week
		diffStr = "1 week"
	case diff < 1*Month:
		diffStr = fmt.Sprintf("%d weeks", diff/Week)
		diff -= diff / Week * Week

	case diff < 2*Month:
		diff -= 1 * Month
		diffStr = "1 month"
	case diff < 1*Year:
		diffStr = fmt.Sprintf("%d months", diff/Month)
		diff -= diff / Month * Month

	case diff < 2*Year:
		diff -= 1 * Year
		diffStr = "1 year"
	default:
		diffStr = fmt.Sprintf("%d years", diff/Year)
		diff = 0
	}
	return diff, diffStr
}

// TimeSincePro calculates the time interval and generate full user-friendly string.
func TimeSincePro(then time.Time) string {
	now := time.Now()
	diff := now.Unix() - then.Unix()

	if then.After(now) {
		return "future"
	}

	var timeStr, diffStr string
	for {
		if diff == 0 {
			break
		}

		diff, diffStr = computeTimeDiff(diff)
		timeStr += ", " + diffStr
	}
	return strings.TrimPrefix(timeStr, ", ")
}
