package tools

/***************************************************************************************
**文件名：Func.go
**包名称：tools
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/

import (
	"strings"
)

//去除空白位
func WipeOutBlank(input []string) []string {
	ret := make([]string, 0, 5)
	for i := 0; i < len(input); i++ {
		if input[i] != "" {
			ret = append(ret, input[i])
		}
	}
	return ret
}

//去除空白前缀
func WipeOutBlankPrefix(input string) string {
	if IsBlank(input) == true {
		return ""
	}
	for i := 0; i < len(input); i++ {
		if input[i] != ' ' && input[i] != '\t' && input[i] != '\n' {
			input = input[i:]
			break
		}
	}
	return input
}

//去除空白字符
func WipeOutBlankString(input string) string {
	input = strings.Replace(input, " ", "", -1)
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "\t", "", -1)
	return input
}

//判断是不是空白字符
func IsBlank(input string) bool {
	if WipeOutBlankString(input) == "" {
		return true
	} else {
		return false
	}
}

//去除注释
func WipeOutNote(input string) string {
	parts := strings.Split(input, "\"")
	var position int = 0
	for i := 0; i < len(parts); i++ {
		if i%2 == 0 {
			index := strings.Index(parts[i], "//")
			if index >= 0 {
				position = position + index
				input = input[:position]
				break
			} else {
				position = position + len(parts[i])
			}
		} else {
			position = position + len(parts[i])
		}
	}
	return input
}
