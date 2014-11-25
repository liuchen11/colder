package parser

/***************************************************************************************
**文件名：Expr.go
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

//Expr 表达式类
type Expr struct {
	expr_left  *Expr  //表达式左面
	expr_right *Opt   //表达式右面
	expr_op    string //表达式符号

	Expr_value_type string //计算出的表达式类型
	Expr_value      string //计算出的表达式值
}

//Expr表达式报错
func ExprError() {
	fmt.Println("there is an error in expr...")
}

//获得一个Expr的指针
func GetExpr(tokenlist []Token) (*Expr, bool) {
	fmt.Print("EXPR : ")
	showTokenlist(tokenlist)

	ret := new(Expr)
	size := len(tokenlist)
	//处理开头为非的情况
	if tokenlist[0].Content == "!" && tokenlist[0].Type == "Operator" {
		ret.expr_op = "!"
		flag := false
		ret.expr_left, flag = GetExpr(tokenlist[1:size])
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			ExprError()
			return nil, false
		}
		//fmt.Println("!")
		//showTokenlist(tokenlist[1:size])
		return ret, true
	} else {
		i := 0
		tmp1 := 0
		tmp2 := 0
		//寻找右面第一个没有嵌套的运算符号
		for i = size - 1; i >= 0; i-- {
			if tmp2 == 0 && tmp1 == 0 && ((tokenlist[i].Type == "Operator" && tokenlist[i].Content == "==") || (tokenlist[i].Type == "Operator" && tokenlist[i].Content == "!=") || (tokenlist[i].Type == "Operator" && tokenlist[i].Content == ">=") || (tokenlist[i].Type == "Operator" && tokenlist[i].Content == "<=") || (tokenlist[i].Type == "Operator" && tokenlist[i].Content == ">") || (tokenlist[i].Type == "Operator" && tokenlist[i].Content == "<") || (tokenlist[i].Type == "Operator" && tokenlist[i].Content == "&&") || (tokenlist[i].Type == "Operator" && tokenlist[i].Content == "||")) {
				break
			} else {
				//处理括号嵌套
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
			//只有其中一只
			ret.expr_op = ""
			flag := false
			ret.expr_right, flag = GetOpt(tokenlist)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				ExprError()
				return nil, false
			}
			//fmt.Println("Only Opt:")
			//showTokenlist(tokenlist)
			return ret, true
		} else {
			//如果有左右两只
			left := tokenlist[0:i]
			right := tokenlist[i+1 : size]
			ret.expr_op = tokenlist[i].Content
			flag := false
			//左面是一个表达式Expr
			ret.expr_left, flag = GetExpr(left)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				ExprError()
				return nil, false
			}
			//右面是一个Opt
			ret.expr_right, flag = GetOpt(right)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				ExprError()
				return nil, false
			}
			//fmt.Print("Left : ")
			//showTokenlist(left)
			//fmt.Print("Right : ")
			//showTokenlist(right)
			return ret, true
		}
	}
}

//调试使用：打印缩进
func PrintIndent(level int) {
	for i := 0; i < level; i++ {
		fmt.Print("\t")
	}
}

//调试使用：打印一个Expr的内容
func (a *Expr) Show(level int) {
	PrintIndent(level)
	fmt.Print("EXPR:")
	if a.expr_op == "" {
		fmt.Println()
		//a.expr_right.Show(level + 1)
	} else {
		if a.expr_op == "!" {
			fmt.Println("!")
			a.expr_left.Show(level + 1)
		} else {
			fmt.Println(a.expr_op)
			a.expr_left.Show(level + 1)
			//a.expr_right.Show(level+1)
		}
	}
}

//Expr运行时候报错
func ExprRunError() {
	fmt.Println("RuntimeError:Error in Expr")
}

//Expr执行
func (a *Expr) Exe(na *NameTable, fu *FuncTable) bool {
	//如果只有一只
	if a.expr_op == "" {
		flag := a.expr_right.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			ExprRunError()
			return false
		}
		//直接进行赋值即可
		a.Expr_value = a.expr_right.Opt_value
		a.Expr_value_type = a.expr_right.Opt_value_type
		return true
	}
	//如果为非
	if a.expr_op == "!" {
		flag := a.expr_left.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			ExprRunError()
			return false
		}
		//要求结果为bool类型
		if a.expr_left.Expr_value_type != "bool" {
			ExprRunError()
			return false
		}
		a.Expr_value_type = "bool"
		if a.expr_left.Expr_value == "false" {
			a.Expr_value = "true"
		} else {
			a.Expr_value = "false"
		}
		return true
	}
	//其他情况，分别对两边使用运算函数
	flag := a.expr_left.Exe(na, fu)
	//判断是否在下面出错，完成栈式报错
	if flag == false {
		ExprRunError()
		return false
	}
	flag = a.expr_right.Exe(na, fu)
	//判断是否在下面出错，完成栈式报错
	if flag == false {
		ExprRunError()
		return false
	}
	//记录两侧的类型和值
	type1 := a.expr_left.Expr_value_type
	type2 := a.expr_right.Opt_value_type
	value1 := a.expr_left.Expr_value
	value2 := a.expr_right.Opt_value

	//对符号进行判断
	if a.expr_op == "&&" || a.expr_op == "||" {
		if type1 != "bool" || type2 != "bool" {
			ExprRunError()
			return false
		}
		if a.expr_op == "&&" {
			//符号为&&
			if value1 == "true" && value2 == "true" {
				a.Expr_value_type = "bool"
				a.Expr_value = "true"
				return true
			} else {
				a.Expr_value_type = "bool"
				a.Expr_value = "false"
				return true
			}
		} else {
			//符号为||
			if value1 == "false" && value2 == "false" {
				a.Expr_value = "false"
				a.Expr_value_type = "bool"
				return true
			} else {
				a.Expr_value = "true"
				a.Expr_value_type = "bool"
				return true
			}
		}
	}
	//如果符号为判断相等或不相等
	if a.expr_op == "==" || a.expr_op == "!=" {
		//两者的类型必须匹配
		if type1 != type2 {
			ExprRunError()
			return false
		}
		//符号进行讨论
		if a.expr_op == "==" {
			//由于使用字符串，所以可以直接进行比对
			if value1 == value2 {
				a.Expr_value = "true"
				a.Expr_value_type = "bool"
				return true
			} else {
				a.Expr_value = "false"
				a.Expr_value_type = "bool"
				return true
			}
		} else {
			//同上
			if value1 == value2 {
				a.Expr_value = "false"
				a.Expr_value_type = "bool"
				return true
			} else {
				a.Expr_value = "true"
				a.Expr_value_type = "bool"
				return true
			}
		}
	}
	//接下里的判断只能在两边都是整数的情况进行
	if type1 != "int" || type2 != "int" {
		ExprRunError()
		return false
	}
	//计算两边的值
	left_value_int, _ := strconv.ParseInt(a.expr_left.Expr_value, 10, 64)
	right_value_int, _ := strconv.ParseInt(a.expr_right.Opt_value, 10, 64)
	//针对运算符进行多分支判断
	switch a.expr_op {
	case ">=":
		//如果为>=
		if left_value_int >= right_value_int {
			a.Expr_value = "true"
			a.Expr_value_type = "bool"
			return true
		} else {
			a.Expr_value = "false"
			a.Expr_value_type = "bool"
			return true
		}
	case "<=":
		//如果为<=
		if left_value_int <= right_value_int {
			a.Expr_value = "true"
			a.Expr_value_type = "bool"
			return true
		} else {
			a.Expr_value = "false"
			a.Expr_value_type = "bool"
			return true
		}
	case ">":
		//如果为>
		if left_value_int > right_value_int {
			a.Expr_value = "true"
			a.Expr_value_type = "bool"
			return true
		} else {
			a.Expr_value = "false"
			a.Expr_value_type = "bool"
			return true
		}
	case "<":
		//如果为<
		if left_value_int < right_value_int {
			a.Expr_value = "true"
			a.Expr_value_type = "bool"
			return true
		} else {
			a.Expr_value = "false"
			a.Expr_value_type = "bool"
			return true
		}
	}
	//如果跳过多分支判断，直接报错
	ExprRunError()
	return false
}
