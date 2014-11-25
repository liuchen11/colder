package parser

/***************************************************************************************
**文件名：Opt.go
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

type Opt struct {
	opt_op    string
	opt_left  *Opt
	opt_right *Term

	Opt_value      string
	Opt_value_type string
}

func OptError() {
	fmt.Println("there is an error in Opt....")
}
func OptRunError() {
	fmt.Println("RunTimeError:Error in Opt....")
}
func GetOpt(tokenlist []Token) (*Opt, bool) {
	fmt.Print("OPT : ")
	showTokenlist(tokenlist)
	ret := new(Opt)
	size := len(tokenlist)
	if tokenlist[0].Content == "-" && tokenlist[0].Type == "Operator" {
		ret.opt_op = "u-"
		flag := false
		ret.opt_left, flag = GetOpt(tokenlist[1:size])
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			OptError()
			return nil, false
		}
		return ret, true
	} else {
		tmp1 := 0
		tmp2 := 0
		i := size - 1
		for i = size - 1; i >= 0; i-- {
			if (tmp2 == 0 && tmp1 == 0 && tokenlist[i].Content == "+" && tokenlist[i].Type == "Operator") || (tmp2 == 0 && tmp1 == 0 && tokenlist[i].Content == "-" && tokenlist[i].Type == "Operator") {
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
		if i == -1 {
			ret.opt_op = ""
			flag := false
			ret.opt_right, flag = GetTerm(tokenlist)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				OptError()
				return nil, false
			}
			return ret, true
		} else {
			left := tokenlist[0:i]
			right := tokenlist[i+1 : size]
			flag := false
			ret.opt_left, flag = GetOpt(left)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				OptError()
				return nil, false
			}
			ret.opt_right, flag = GetTerm(right)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				OptError()
				return nil, false
			}
			if tokenlist[i].Content == "-" {
				ret.opt_op = "b-"
			} else {
				ret.opt_op = "+"
			}
			return ret, true
		}
	}
	return ret, true
}
func (a *Opt) Exe(na *NameTable, fu *FuncTable) bool {
	if a.opt_op == "" {
		flag := a.opt_right.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			OptRunError()
			return false
		}
		a.Opt_value = a.opt_right.Term_value
		a.Opt_value_type = a.opt_right.Term_value_type
		return true
	}
	if a.opt_op == "u-" {
		flag := a.opt_left.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			OptRunError()
			return false
		}
		type1 := a.opt_left.Opt_value_type
		if type1 != "int" {
			OptRunError()
			return false
		}
		a.Opt_value_type = "int"
		left_value_int, _ := strconv.ParseInt(a.opt_left.Opt_value, 10, 64)
		a.Opt_value = strconv.FormatInt(int64(-left_value_int), 10)
		return true
	}
	if a.opt_op == "b-" {
		flag := a.opt_left.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			OptRunError()
			return false
		}
		flag = a.opt_right.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			OptRunError()
			return false
		}
		type1 := a.opt_left.Opt_value_type
		type2 := a.opt_right.Term_value_type
		if type1 != "int" || type2 != "int" {
			OptRunError()
			return false
		}
		left_value_int, _ := strconv.ParseInt(a.opt_left.Opt_value, 10, 64)
		right_value_int, _ := strconv.ParseInt(a.opt_right.Term_value, 10, 64)
		a.Opt_value = strconv.FormatInt(int64(left_value_int-right_value_int), 10)
		a.Opt_value_type = "int"
		return true
	}
	if a.opt_op == "+" {
		flag := a.opt_left.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			OptRunError()
			return false
		}
		flag = a.opt_right.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			OptRunError()
			return false
		}
		type1 := a.opt_left.Opt_value_type
		type2 := a.opt_right.Term_value_type
		if type1 != type2 {
			OptRunError()
			return false
		}
		if type1 == "bool" {
			OptRunError()
			return false
		}
		if type1 == "int" {
			left_value_int, _ := strconv.ParseInt(a.opt_left.Opt_value, 10, 64)
			right_value_int, _ := strconv.ParseInt(a.opt_right.Term_value, 10, 64)
			a.Opt_value = strconv.FormatInt(int64(left_value_int+right_value_int), 10)
			a.Opt_value_type = "int"
			return true
		} else {
			if type1 == "string" {
				a.Opt_value_type = "string"
				a.Opt_value = a.opt_left.Opt_value + a.opt_right.Term_value
				return true
			} else {
				OptRunError()
				return false
			}
		}
	}
	return false
}
