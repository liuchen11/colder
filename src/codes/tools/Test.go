package tools

/***************************************************************************************
**文件名：Test.go
**包名称：tools
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


import (
	"bufio"
	"io"
	"os"
	"strings"
)

//比较两个文件是否相同
func CompareTwoFiles(file1 string, file2 string, strict bool) (int, bool) {
	fin1, err1 := os.Open(file1)
	fin2, err2 := os.Open(file2)
	defer fin1.Close()
	defer fin2.Close()
	if err1 != nil || err2 != nil {
		return 1, false
	}
	reader1 := bufio.NewReader(fin1)
	reader2 := bufio.NewReader(fin2)
	for {
		buf1, _, error1 := reader1.ReadLine()
		buf2, _, error2 := reader2.ReadLine()
		switch {
		case error1 == io.EOF && error2 == io.EOF:
			return 0, true
		case error1 == io.EOF || error2 == io.EOF:
			return 0, false
		case error1 != nil || error2 != nil:
			return 2, false
		default:
			buffer1 := string(buf1)
			buffer2 := string(buf2)
			if strict == false {
				buffer1 = strings.Replace(buffer1, " ", "", -1)
				buffer1 = strings.Replace(buffer1, "\t", "", -1)
				buffer2 = strings.Replace(buffer2, " ", "", -1)
				buffer2 = strings.Replace(buffer2, "\t", "", -1)
			}
			if buffer1 != buffer2 {
				return 0, false
			}
		}
	}
}

