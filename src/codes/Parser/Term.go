package parser

/***************************************************************************************
**文件名：Term.go
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

type Term struct {
	term_left  *Term
	term_op    string
	term_right *Factor

	Term_value_type string
	Term_value      string
}

func TermError() {
	fmt.Println("there is an error in term....")
}
func TermRunError() {
	fmt.Println("RUNTIMEERROR:TERM.....")
}

func GetTerm(tokenlist []Token) (*Term, bool) {
	fmt.Print("TERM : ")
	showTokenlist(tokenlist)
	ret := new(Term)
	size := len(tokenlist)
	tmp1 := 0
	tmp2 := 0
	i := size - 1
	for i = size - 1; i >= 0; i-- {
		if (tmp2 == 0 && tmp1 == 0 && tokenlist[i].Content == "*" && tokenlist[i].Type == "Operator") || (tmp2 == 0 && tmp1 == 0 && tokenlist[i].Content == "/" && tokenlist[i].Type == "Operator") {
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
		ret.term_op = ""
		flag := false
		ret.term_right, flag = GetFactor(tokenlist)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			TermError()
			return nil, false
		}
		return ret, true
	} else {
		left := tokenlist[0:i]
		right := tokenlist[i+1 : size]
		flag := false
		ret.term_left, flag = GetTerm(left)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			TermError()
			return nil, false
		}
		ret.term_right, flag = GetFactor(right)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			TermError()
			return nil, false
		}
		ret.term_op = tokenlist[i].Content
		return ret, true
	}
	return ret, true
}

func (a *Term) Exe(name *NameTable, fu *FuncTable) bool {
	if a.term_op == "" {
		flag := a.term_right.Exe(name, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			TermRunError()
			return false
		}
		a.Term_value = a.term_right.Factor_value_single
		a.Term_value_type = a.term_right.Factor_value_type
		return true
	}
	flag := false
	flag = a.term_left.Exe(name, fu)
	//判断是否在下面出错，完成栈式报错
	if flag == false {
		fmt.Println("Error in Term.....")
		return false
	}
	flag = a.term_right.Exe(name, fu)
	//判断是否在下面出错，完成栈式报错
	if flag == false {
		fmt.Println("Error in Term.....")
		return false
	}
	switch a.term_op {
	case "*":
		if a.term_left.Term_value_type == "int" && a.term_right.Factor_value_type == "int" {
			a.Term_value_type = "int"
			left_value_int, _ := strconv.ParseInt(a.term_left.Term_value, 10, 64)
			right_value_int, _ := strconv.ParseInt(a.term_right.Factor_value_single, 10, 64)
			a.Term_value = strconv.FormatInt(int64(left_value_int*right_value_int), 10)
			return true
		}
		fmt.Println("Error in Term...")
		return false
	case "/":
		if a.term_left.Term_value_type == "int" && a.term_right.Factor_value_type == "int" {
			a.Term_value_type = "int"
			left_value_int, _ := strconv.ParseInt(a.term_left.Term_value, 10, 64)
			right_value_int, _ := strconv.ParseInt(a.term_right.Factor_value_single, 10, 64)
			if right_value_int == 0 {
				a.Term_value_type = ""
				fmt.Println("Error in Term.....")
				return false
			}
			a.Term_value = strconv.FormatInt(int64(left_value_int/right_value_int), 10)
			return true
		}
	}
	return false
}
