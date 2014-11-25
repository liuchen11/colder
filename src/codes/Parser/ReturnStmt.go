package parser

/***************************************************************************************
**文件名：ReturnStmt.go
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

type ReturnStmt struct {
	returnstmt_epxr *Expr
}

func ReturnStmtError() {
	fmt.Println("There is an error in ReturnStmt")
}

func GetReturnStmt(tokenlist []Token) (*ReturnStmt, bool) {
	ret := new(ReturnStmt)
	size := len(tokenlist)
	if tokenlist[0].Content == "return" && tokenlist[0].Type == "Keyword" {
		flag := false
		ret.returnstmt_epxr, flag = GetExpr(tokenlist[1:size])
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			ReturnStmtError()
			return nil, false
		}
		return ret, true
	} else {
		ReturnStmtError()
		return nil, false
	}
	return ret, true
}
func ReturnStmtRunError() {
	fmt.Println("RunTimeError:There is an error in ReturnStmt")
}
func (a *ReturnStmt) Exe(na *NameTable, fu *FuncTable) bool {
	flag := a.returnstmt_epxr.Exe(na, fu)
	//判断是否在下面出错，完成栈式报错
	if flag == false {
		ReturnStmtRunError()
		return false
	}
	v_type := a.returnstmt_epxr.Expr_value_type
	v_value := a.returnstmt_epxr.Expr_value
	fu.ReturnFunc(v_type, v_value)
	return true
}
