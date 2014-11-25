package parser

/***************************************************************************************
**文件名：LValue.go
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

type LValue struct {
	lvalue_type       string
	lvalue_token      *Token
	lvalue_expr_right *Expr

	LValue_type  string
	LValue_name  string
	LValue_index int64
}

func LValueError() {
	fmt.Println("There is an error in Lvalue....")
}

func GetLValue(tokenlist []Token) (*LValue, bool) {
	fmt.Print("LVALUE : ")
	showTokenlist(tokenlist)
	ret := new(LValue)
	size := len(tokenlist)
	if size == 1 {
		if tokenlist[0].Type == "Identifier" {
			ret.lvalue_token = copy_token(&tokenlist[0])
			ret.lvalue_type = "ident"
			return ret, true
		} else {
			LValueError()
			return nil, false
		}
	} else {
		if tokenlist[size-1].Content == "]" && tokenlist[size-1].Type == "Operator" {
			tmp1 := 0
			tmp2 := 0
			i := size - 2
			for i = size - 2; i >= 0; i-- {
				if tmp2 == 0 && tmp1 == 0 && tokenlist[i].Content == "[" && tokenlist[i].Type == "Operator" {
					break
				} else {
					if tokenlist[i].Type == "Operator" {
						switch tokenlist[i].Content {
						case "[":
							tmp1--
						case "]":
							tmp1++
						case "(":
							tmp2--
						case ")":
							tmp2++
						}
					}
				}
			}
			right := tokenlist[i+1 : size-1]
			flag := false
			//ret.lvalue_expr_left, flag = GetExpr(left)
			////判断是否在下面出错，完成栈式报错
			//if flag == false {
			//	LValueError()
			//	return nil, false
			//}
			if i != 1 {
				LValueError()
				return nil, false
			}
			if tokenlist[0].Type != "Identifier" {
				LValueError()
				return nil, false
			}
			ret.lvalue_token = copy_token(&tokenlist[i-1])

			ret.lvalue_expr_right, flag = GetExpr(right)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				LValueError()
				return nil, false
			}
			ret.lvalue_type = "expr"
			return ret, true
		} else {
			LValueError()
			return nil, false
		}
	}
	return ret, true
}

func LValueRunError() {
	fmt.Println("RunTimeError:There is an error in LValue")
}
func (a *LValue) Exe(na *NameTable, fu *FuncTable) bool {
	if a.lvalue_type == "expr" {
		flag := a.lvalue_expr_right.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			LValueRunError()
			return false
		}
		if a.lvalue_expr_right.Expr_value_type != "int" {
			LValueRunError()
			return false
		}
		left_value_int, _ := strconv.ParseInt(a.lvalue_expr_right.Expr_value, 10, 64)
		a.LValue_index = left_value_int
		a.LValue_name = a.lvalue_token.Content
		a.LValue_type = "array"
		return true
	} else {
		a.LValue_type = "iden"
		a.LValue_name = a.lvalue_token.Content
		typeit := na.GetType(a.LValue_name)
		switch typeit {
		case "intarr":
			a.LValue_index = int64(na.GetIntlen(a.LValue_name))
		case "stringarr":
			a.LValue_index = int64(na.GetStringlen(a.LValue_name))
		case "boolarr":
			a.LValue_index = int64(na.GetBoollen(a.LValue_name))
		}
		return true
	}
}
func (a *LValue) Show() {
	fmt.Println(a.LValue_type)
	fmt.Println(a.LValue_name)
	if a.LValue_type == "array" {
		fmt.Println(a.LValue_index)
	}
}
