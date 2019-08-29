package log

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	INFO_COLOR = "%c[1;32m%s%c[0m"
	ERROR_COLOR = "%c[1;31m%s%c[0m"
	WARN_COLOR = "%c[1;33m%s%c[0m"
)

var colorMap = map[int]string{
	31: ERROR_COLOR,
	32: INFO_COLOR,
	33: WARN_COLOR,
}

/**
	一般级别
 */
func Info(args ...interface{}){
	if _, file, line, ok := runtime.Caller(1); ok {
		_, filename := path.Split(file)
		prefix := formatLog("INFO", filename, line)
		commonLogger(32, prefix, args)
	} else {
		panic(errors.New("runtime error"))
	}
}

/**
	错误级别
*/
func Error(args ...interface{}){
	if _, file, line, ok := runtime.Caller(1); ok {
		_, filename := path.Split(file)
		prefix := formatLog("ERROR", filename, line)
		commonLogger(31, prefix, args)
	} else {
		panic(errors.New("runtime error"))
	}
}

/**
	警告级别
*/
func Warn(args ...interface{}){
	if _, file, line, ok := runtime.Caller(1); ok {
		_, filename := path.Split(file)
		prefix := formatLog("WARN", filename, line)
		commonLogger(33, prefix, args)
	} else {
		panic(errors.New("runtime error"))
	}
}

/**
	公共的日志生成逻辑
 */
func commonLogger(levelColor int, prefix string, args ...interface{}) {
	switch len(args[0].([]interface{})) {
	case 1:
		combineLogPrefix(colorMap[levelColor], prefix, args[0].([]interface{})[0])
	default:
		typeOf := reflect.TypeOf(args[0].([]interface{})[0])
		if typeOf.Kind() != reflect.String {
			combineLogPrefix(ERROR_COLOR, prefix, "ErrorType, should be a String params at 1 position")
			os.Exit(1)
		}
		f := args[0].([]interface{})[0].(string)
		if !strings.Contains(f, "%") ||
			(strings.Count(f, "%") != len(args[0].([]interface{})[1:]) && !strings.Contains(f, "[")) {
			combineLogPrefix(ERROR_COLOR, prefix, "format error")
			os.Exit(1)
		}

		var msgSlice = args[0].([]interface{})[1:]
		var msgs string
		for _, item := range msgSlice {
			msgs += item.(string)
		}
		if msgs != "" {
			if ok, err := regexp.MatchString(`(\s\S)*(%\[\d+\])(\s\S)*`, f); err == nil && ok {
				re, _ := regexp.Compile(`%\[(\d+)\]`)
				// 为了实现%[1]s的功能，原因是颜色占用了一个位置，需要往前挪一位
				newStr := re.ReplaceAllFunc([]byte(f), func(b1 []byte) []byte {
					reChild, _ := regexp.Compile(`\d+`)
					return reChild.ReplaceAllFunc(b1, func(b2 []byte) []byte{
						num, _ := strconv.Atoi(string(b2))
						return []byte(strconv.Itoa(num + 1))
					})
				})
				fmt.Printf(colorMap[levelColor], 0x1B, prefix, 0x1B)
				fmt.Printf(combineColorFormat(levelColor, string(newStr)), 0x1B, msgs , 0x1B)
				return
			}
			fmt.Printf(colorMap[levelColor], 0x1B, prefix, 0x1B)
			fmt.Printf(combineColorFormat(levelColor, f), 0x1B, msgs , 0x1B)
		} else {
			combineLogPrefix(ERROR_COLOR, prefix, "not found msg")
		}
	}
}

/**
	拼接前缀
 */
func combineLogPrefix(format string, prefix string, msg interface{}) {
	fmt.Printf(format, 0x1B, prefix, 0x1B)
	fmt.Printf(format + "\n", 0x1B, msg, 0x1B)
}

/**
	生成带颜色的格式
*/
func combineColorFormat(levelColor int, s string) string{
	return "%c[1;" + strconv.Itoa(levelColor) + "m"+ s +"%c[0m\n"
}

/**
	日志前缀
*/
func formatLog(level, filename string, line int) string{
	return "[" + level + "]|" + filename + ":" + strconv.Itoa(line) + "|" + time.Now().String()[:19] + "| "
}
