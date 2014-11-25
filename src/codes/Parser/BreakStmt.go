package parser

/***************************************************************************************
**文件名：BreakStmt.go
**包名称：parser
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


import (
	"fmt"
)

type BreakStmt struct {
}

func BreakStmtError() {
	fmt.Println("there is an error in BreakStmt....")
}

func GetBreakStmt(tokenlist []Token) (*BreakStmt, bool) {
	fmt.Print("BreakStmt : ")
	showTokenlist(tokenlist)
	size := len(tokenlist)
	if size != 1 {
		BreakStmtError()
		return nil, false
	} else {
		if tokenlist[0].Content != "break" || tokenlist[0].Type != "Keyword" {
			BreakStmtError()
			return nil, false
		}
	}
	ret := new(BreakStmt)
	return ret, true
}

