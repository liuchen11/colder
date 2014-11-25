package parser

/***************************************************************************************
**文件名：Stmt.go
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

type Stmt struct {
	stmt_type   string
	stmt_var    *VarStmt
	stmt_simple *SimpleStmt
	stmt_switch *SwitchStmt
	stmt_for    *ForStmt
	stmt_break  *BreakStmt
	stmt_block  *BlockStmt
	stmt_return *ReturnStmt
}

func StmtError() {
	fmt.Println("There is an error in Stmt...")
}

func GetStmt(tokenlist []Token) (*Stmt, bool) {

	fmt.Print("STMT:")
	showTokenlist(tokenlist)
	ret := new(Stmt)
	size := len(tokenlist)
	if size == 0 {
		StmtError()
		return nil, false
	}
	if size == 1 {
		if tokenlist[0].Content == ";" && tokenlist[0].Type == "Operator" {
			ret.stmt_type = "simple"
			flag := false
			ret.stmt_simple, flag = GetSimpleStmt(tokenlist[0 : size-1])
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				StmtError()
				return nil, false
			}
			return ret, true
		}
	}

	//ReturnStmt
	if tokenlist[0].Type == "Keyword" && tokenlist[0].Content == "return" {
		if tokenlist[size-1].Content != ";" || tokenlist[size-1].Type != "Operator" {
			StmtError()
			return nil, false
		}
		ret.stmt_type = "return"
		flag := false
		ret.stmt_return, flag = GetReturnStmt(tokenlist[0 : size-1])
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			StmtError()
			return nil, false
		}
		return ret, true
	}

	//SimpleStmt, BreakStmt, VarStmt
	if tokenlist[0].Type == "Identifier" || tokenlist[0].Type == "Reservedword" {
		//SimpleStmt
		ret.stmt_type = "simple"
		flag := false
		ret.stmt_simple, flag = GetSimpleStmt(tokenlist[0 : size-1])
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			StmtError()
			return nil, false
		}
		return ret, true
	}
	if (tokenlist[0].Type == "Keyword") && (tokenlist[0].Content == "int" || tokenlist[0].Content == "bool" || tokenlist[0].Content == "string") {
		//VarStmt
		ret.stmt_type = "var"
		flag := false
		ret.stmt_var, flag = GetVarStmt(tokenlist[0 : size-1])
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			StmtError()
			return nil, false
		}
		return ret, true
	}
	if size == 2 && tokenlist[0].Content == "break" && tokenlist[0].Type == "Keyword" {
		//BreakStmt
		ret.stmt_type = "break"
		flag := false
		ret.stmt_break, flag = GetBreakStmt(tokenlist[0 : size-1])
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			StmtError()
			return nil, false
		}
		return ret, true
	}

	//SwitchStmt, ForStmt, BlockStmt
	if tokenlist[0].Type == "Keyword" {
		if tokenlist[0].Content == "for" {
			ret.stmt_type = "for"
			flag := false
			ret.stmt_for, flag = GetForStmt(tokenlist)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				StmtError()
				return nil, false
			}
			return ret, true
		} else {
			if tokenlist[0].Content == "switch" {
				ret.stmt_type = "switch"
				flag := false
				ret.stmt_switch, flag = GetSwitchStmt(tokenlist)
				//判断是否在下面出错，完成栈式报错
				if flag == false {
					StmtError()
					return nil, false
				}
				return ret, true
			}
			StmtError()
			return nil, false
		}
	} else {
		//Block
		fmt.Println(tokenlist)
		if tokenlist[0].Content == "{" && tokenlist[0].Type == "Operator" && tokenlist[size-1].Type == "Operator" && tokenlist[size-1].Content == "}" {
			ret.stmt_type = "block"
			flag := false
			ret.stmt_block, flag = GetBlockStmt(tokenlist)

			//判断是否在下面出错，完成栈式报错
			if flag == false {
				StmtError()
				return nil, false
			}
			return ret, true
		} else {
			StmtError()
			return nil, false
		}
	}
	return nil, false
}
func StmtRunError() {
	fmt.Println("RunTimeError : There is an error in Stmt...")
}
func (a *Stmt) Exe(na *NameTable, fu *FuncTable) bool {
	switch a.stmt_type {
	case "block":
		flag := a.stmt_block.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			StmtRunError()
			return false
		}
		return true
	case "switch":
		flag := a.stmt_switch.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			StmtRunError()
			return false
		}
		return true
	case "for":
		flag := a.stmt_for.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			StmtRunError()
			return false
		}
		return true
	case "return":
		flag := a.stmt_return.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			StmtError()
			return false
		}
		return true
	case "var":
		flag := a.stmt_var.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			StmtRunError()
			return false
		}
		return true
	case "simple":
		flag := a.stmt_simple.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			StmtRunError()
			return false
		}
		return true
	default:
		StmtRunError()
		return false
	}
	return false
}
