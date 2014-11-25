package parser

/***************************************************************************************
**文件名：SwitchStmt.go
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

type SwitchStmt struct {
	switchstmt_condition []*BoolExpr
	switchstmt_content   []*Stmt
}

func SwitchStmtError() {
	fmt.Println("There is an error in SwitchStmt....")
}

func GetSwitchStmt(tokenlist []Token) (*SwitchStmt, bool) {
	fmt.Print("SWITCHSTMT:")
	showTokenlist(tokenlist)
	ret := new(SwitchStmt)
	ret.switchstmt_condition = make([]*BoolExpr, 0, 5)
	ret.switchstmt_content = make([]*Stmt, 0, 5)
	size := len(tokenlist)
	if size < 7 {
		SwitchStmtError()
		return nil, false
	}

	if tokenlist[0].Content != "switch" || tokenlist[0].Type != "Keyword" {
		SwitchStmtError()
		return nil, false
	}
	i := 1
	if tokenlist[i].Content != "{" || tokenlist[i].Type != "Operator" {
		SwitchStmtError()
		return nil, false
	}
	i++

	for {
		if i >= size {

			SwitchStmtError()
			return nil, false
		}
		if tokenlist[i].Type == "Keyword" && tokenlist[i].Content == "default" {
			if tokenlist[size-1].Content == "}" && tokenlist[size-1].Type == "Operator" {
				tmp := tokenlist[i+2 : size-1]
				next, flag := GetStmt(tmp)
				//判断是否在下面出错，完成栈式报错
				if flag == false {
					SwitchStmtError()
					return nil, false
				}
				ret.switchstmt_content = append(ret.switchstmt_content, next)
				return ret, true
			} else {
				SwitchStmtError()
				return nil, false
			}
		} else {
			if tokenlist[i].Type == "Keyword" && tokenlist[i].Content == "case" {
				//Find Next condition
				j := i + 1
				if tokenlist[j].Content != "(" || tokenlist[j].Type != "Operator" {
					SwitchStmtError()
					return nil, false
				}
				for j = i + 1; j < size; j++ {
					if tokenlist[j].Content == ":" && tokenlist[j].Type == "Operator" {
						break
					}
				}
				if j == size {
					SwitchStmtError()
					return nil, false
				}
				if tokenlist[j-1].Content != ")" || tokenlist[j-1].Type != "Operator" {
					SwitchStmtError()
					return nil, false
				}
				tmp := tokenlist[i+2 : j-1]
				next_condition, flag_condition := GetBoolExpr(tmp)
				if flag_condition == false {
					SwitchStmtError()
					return nil, false
				}
				ret.switchstmt_condition = append(ret.switchstmt_condition, next_condition)
				i = j
				cnt := 0
				for j = i + 1; j < size; j++ {
					if cnt == 0 && ((tokenlist[j].Content == "case" && tokenlist[j].Type == "Keyword") || (tokenlist[j].Content == "default" && tokenlist[j].Type == "Keyword")) {
						break
					} else {
						if tokenlist[j].Content == "{" && tokenlist[j].Type == "Operator" {
							cnt--
						} else {
							if tokenlist[j].Content == "}" && tokenlist[j].Type == "Operator" {
								cnt++
							}
						}
					}
				}
				tmp = tokenlist[i+1 : j]
				next_stmt, flag_stmt := GetStmt(tmp)
				if flag_stmt == false {
					SwitchStmtError()
					return nil, false
				}
				ret.switchstmt_content = append(ret.switchstmt_content, next_stmt)
				i = j
			} else {
				SwitchStmtError()
				return nil, false
			}
		}
	}
	return ret, true
}

func SwitchStmtRunError() {
	fmt.Println("RunTimeError:There is an error in SwitchStmt...")
}

func (a *SwitchStmt) Exe(na *NameTable, fu *FuncTable) bool {
	size := len(a.switchstmt_condition)
	size1 := len(a.switchstmt_content)
	if size1-size != 1 {
		SwitchStmtRunError()
		return false
	}
	flag := false
	for i := 0; i < size; i++ {
		flag = a.switchstmt_condition[i].Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			SwitchStmtRunError()
			return false
		}
		if a.switchstmt_condition[i].BoolExpr_value == "true" {
			flag = a.switchstmt_content[i].Exe(na, fu)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				SwitchStmtRunError()
				return false
			}
			return true
		}
	}
	flag = a.switchstmt_content[size1-1].Exe(na, fu)
	//判断是否在下面出错，完成栈式报错
	if flag == false {
		SwitchStmtRunError()
		return false
	}
	return true
}
