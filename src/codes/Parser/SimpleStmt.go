package parser

/***************************************************************************************
**文件名：SimpleStmt.go
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
	"strings"
)

type SimpleStmt struct {
	simplestmt_type   string
	simplestmt_call   *Call
	simplestmt_expr   *Expr
	simplestmt_lvalue *LValue
}

func SimpleStmtError() {
	fmt.Println("There is an error in SimplerStmt....")
}
func SimpleStmtRunError() {
	fmt.Println("RunTimeError:Error in SimpleStmt....")
}

func GetSimpleStmt(tokenlist []Token) (*SimpleStmt, bool) {
	fmt.Print("SimpleStmt:")
	showTokenlist(tokenlist)
	ret := new(SimpleStmt)
	size := len(tokenlist)
	if size == 0 {
		ret.simplestmt_type = "empty"
		return ret, true
	} else {
		i := 0
		for i = 0; i < size; i++ {
			if tokenlist[i].Content == "=" && tokenlist[i].Type == "Operator" {
				break
			}
		}
		if i == size {
			//没有找到
			ret.simplestmt_type = "call"
			flag := false
			ret.simplestmt_call, flag = GetCall(tokenlist)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				SimpleStmtError()
				return nil, false
			}
			return ret, true
		} else {
			//找到
			left := tokenlist[:i]
			right := tokenlist[i+1 : size]
			flag := false
			ret.simplestmt_lvalue, flag = GetLValue(left)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				SimpleStmtError()
				return nil, false
			}
			ret.simplestmt_expr, flag = GetExpr(right)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				SimpleStmtError()
				return nil, false
			}
			ret.simplestmt_type = "lvalue"
			return ret, true
		}
	}
	return ret, true
}
func (a *SimpleStmt) Exe(na *NameTable, fu *FuncTable) bool {
	switch a.simplestmt_type {
	case "empty":
		return true
	case "call":
		//TODOIT...
	case "lvalue":
		flag := a.simplestmt_lvalue.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			SimpleStmtRunError()
			return false
		}

		flag = a.simplestmt_expr.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			SimpleStmtRunError()
			return false
		}

		if a.simplestmt_lvalue.LValue_type == "iden" {
			//single
			left_type := na.GetType(a.simplestmt_lvalue.LValue_name)
			right_type := a.simplestmt_expr.Expr_value_type
			if left_type == right_type {
				switch left_type {
				case "int":
					v_name := a.simplestmt_lvalue.LValue_name
					v_value, _ := strconv.ParseInt(a.simplestmt_expr.Expr_value, 10, 32)
					na.SetInt(v_name, int(v_value))
					return true
				case "string":
					v_name := a.simplestmt_lvalue.LValue_name
					v_value := a.simplestmt_expr.Expr_value
					na.SetString(v_name, v_value)
					return true
				case "bool":
					v_name := a.simplestmt_lvalue.LValue_name
					v_value := false
					if a.simplestmt_expr.Expr_value == "true" {
						v_value = true
					} else {
						v_value = false
					}
					na.SetBool(v_name, v_value)
					return true
				case "intarr":
					tmp := strings.Split(a.simplestmt_expr.Expr_value, ",")
					l_len := int(a.simplestmt_lvalue.LValue_index)
					r_len := len(tmp)
					if l_len != r_len {
						SimpleStmtRunError()
						return false
					}
					v_name := a.simplestmt_lvalue.LValue_name
					for i := 0; i < l_len; i++ {
						v_value, _ := strconv.ParseInt(tmp[i], 10, 32)
						na.SetIntArr(v_name, i, int(v_value))
					}
					return true
				case "boolarr":
					tmp := strings.Split(a.simplestmt_expr.Expr_value, ",")
					l_len := int(a.simplestmt_lvalue.LValue_index)
					r_len := len(tmp)
					if l_len != r_len {
						SimpleStmtRunError()
						return false
					}
					v_name := a.simplestmt_lvalue.LValue_name
					for i := 0; i < l_len; i++ {
						v_value := false
						if tmp[i] == "false" {
							v_value = false
						} else {
							v_value = true
						}
						na.SetBoolArr(v_name, i, v_value)
					}
					return true
				case "stringarr":
					tmp := strings.Split(a.simplestmt_expr.Expr_value, ",")
					l_len := int(a.simplestmt_lvalue.LValue_index)
					r_len := len(tmp)
					if l_len != r_len {
						SimpleStmtRunError()
						return false
					}
					v_name := a.simplestmt_lvalue.LValue_name
					for i := 0; i < l_len; i++ {
						v_value := tmp[i]
						na.SetStringArr(v_name, i, v_value)
					}
					return true
				default:
					SimpleStmtRunError()
					return false
				}
			} else {
				SimpleStmtRunError()
				return false
			}
		} else {
			//array
			//!!!!TODO!!!!!
			label := int(a.simplestmt_lvalue.LValue_index)
			v_name := a.simplestmt_lvalue.LValue_name
			right_type := a.simplestmt_expr.Expr_value_type
			switch right_type {
			case "int":
				if na.GetType(v_name) != "intarr" {
					SimpleStmtRunError()
					return false
				}
				v_value, _ := strconv.ParseInt(a.simplestmt_expr.Expr_value, 10, 64)
				na.SetIntArr(v_name, label, int(v_value))
				return true
			case "bool":
				if na.GetType(v_name) != "boolarr" {
					SimpleStmtRunError()
					return false
				}
				value := false
				if a.simplestmt_expr.Expr_value == "true" {
					value = true
				}
				na.SetBoolArr(v_name, label, value)
				return true
			case "string":
				if na.GetType(v_name) != "stringarr" {
					SimpleStmtRunError()
					return false
				}
				na.SetStringArr(v_name, label, a.simplestmt_expr.Expr_value)
				return true
			default:
				SimpleStmtRunError()
				return false
			}
		}
	default:
		SimpleStmtRunError()
		return false
	}
	return false
}
