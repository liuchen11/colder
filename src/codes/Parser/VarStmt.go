package parser

/***************************************************************************************
**文件名：VarStmt.go
**包名称：parser
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/

import (
	"fmt"
	"strconv"
)

type VarStmt struct {
	varstmt_typedef *TypeDef
	varstmt_token   *Token
	varstmt_length  int64
	varstmt_type    string
}

func VarStmtError() {
	fmt.Println("There is an error in varstmt....")
}
func VarStmtRunError() {
	fmt.Println("RunTimeError:In VarStmt....")
}

func GetVarStmt(tokenlist []Token) (*VarStmt, bool) {
	fmt.Print("VARSTMT:")
	showTokenlist(tokenlist)
	ret := new(VarStmt)
	size := len(tokenlist)
	if size <= 1 {
		VarStmtError()
		return nil, false
	}
	if tokenlist[size-1].Type == "Identifier" {
		ret.varstmt_type = "single"
		ret.varstmt_token = copy_token(&tokenlist[size-1])
		flag := false
		ret.varstmt_typedef, flag = GetTypeDef(tokenlist[:size-1])
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			VarStmtError()
			return nil, false
		} else {
			return ret, true
		}
	} else {
		if size >= 4 && tokenlist[size-1].Content == "]" && tokenlist[size-1].Type == "Operator" && tokenlist[size-2].Type == "Int" && tokenlist[size-3].Type == "Operator" && tokenlist[size-3].Content == "[" && tokenlist[size-4].Type == "Identifier" {
			ret.varstmt_type = "array"
			ret.varstmt_token = copy_token(&tokenlist[size-4])
			flag := false
			ret.varstmt_typedef, flag = GetTypeDef(tokenlist[0 : size-4])
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				VarStmtError()
				return nil, false
			}
			ret.varstmt_length, _ = strconv.ParseInt(tokenlist[size-2].Content, 10, 64)
			return ret, true
		} else {
			VarStmtError()
			return nil, false
		}
	}
	return ret, true
}

func (a *VarStmt) Exe(na *NameTable, fu *FuncTable) bool {
	if a.varstmt_type == "single" {
		//For Single
		flag := a.varstmt_typedef.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			VarStmtRunError()
			return false
		}
		v_name := a.varstmt_token.Content
		v_type := a.varstmt_typedef.Typedef_type
		na.AddVariable(v_type, v_name)
		return true
	} else {
		//For Array
		flag := a.varstmt_typedef.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			VarStmtRunError()
			return false
		}
		v_name := a.varstmt_token.Content
		v_type := a.varstmt_typedef.Typedef_type
		v_len := int(a.varstmt_length)
		na.AddArray(v_type+"arr", v_name, v_len)
		return true
	}
}
