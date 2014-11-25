package parser

/***************************************************************************************
**文件名：Call.go
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

//类Call
type Call struct {
	call_name *Token  //函数名称
	call_args []*Expr //调用参数列表

	Call_type  string //调用返回的类型
	Call_value string //调用返回的值
}

//Call进行报错处理
func CallError() {
	fmt.Println("There is an error in Call")
}

//显示一个Token数组
func showTokenlist(tokenlist []Token) {
	fmt.Print("{")
	size := len(tokenlist)
	for i := 0; i < size; i++ {
		fmt.Print(tokenlist[i].Content + " ")
	}
	fmt.Println("}")
}

//获得Call指针
func GetCall(tokenlist []Token) (*Call, bool) {
	fmt.Print("CALL : ")
	showTokenlist(tokenlist)
	ret := new(Call)
	//第一个Token的COntent作为调用的名称
	ret.call_name = copy_token(&tokenlist[0])

	//获取参数列表
	size := len(tokenlist)
	ret.call_args = make([]*Expr, 0, 5)
	if tokenlist[1].Content == "(" && tokenlist[1].Type == "Operator" && tokenlist[size-1].Content == ")" && tokenlist[size-1].Type == "Operator" {
		now := make([]Token, 0, 5)
		tmp1 := 0
		tmp2 := 0

		//每次寻找第一个不进入括号嵌套的逗号，利用逗号进行分割

		for i := 2; i < size-1; i++ {
			if tmp2 == 0 && tmp1 == 0 && tokenlist[i].Content == "," && tokenlist[i].Type == "Operator" {
				//showTokenlist(now)
				//获得参数列表中的表达式
				nextexpr, flag := GetExpr(now)
				//判断是否在下面出错，完成栈式报错
				if flag == false {
					CallError()
					return nil, false
				}
				//刷新数组，进一步获得Token
				now = make([]Token, 0, 5)
				ret.call_args = append(ret.call_args, nextexpr)
			} else {

				//如果是括号，进行嵌套统计
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
				now = append(now, tokenlist[i])
			}
		}

		//单独处理最后一个表达式
		if len(now) != 0 {
			//showTokenlist(now)
			nextexpr, flag := GetExpr(now)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				FactorError()
				return nil, false
			}
			ret.call_args = append(ret.call_args, nextexpr)
		}
	} else {
		CallError()
		return ret, false
	}
	return ret, true
}

//Call报错，运行时错误
func CallRunError() {
	fmt.Println("RunTimeError:There is an error in Call...")
}

//Call 运行函数
func (a *Call) Exe(na *NameTable, fu *FuncTable) bool {
	name := a.call_name.Content
	size := len(a.call_args)
	now_type := make([]string, 0, 5)
	now_value := make([]string, 0, 5)

	//参数运算
	for i := 0; i < size; i++ {
		flag := a.call_args[i].Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			CallRunError()
			return false
		}
		now_type = append(now_type, a.call_args[i].Expr_value_type)
		now_value = append(now_value, a.call_args[i].Expr_value)
	}

	//调用函数
	flag1 := fu.Exe(name, na, now_type, now_value)
	if flag1 == false {
		CallRunError()
		return false
	}

	//进行赋值
	a.Call_type, a.Call_value = fu.Get(name)
	//fmt.Println("BOOMB: " + a.Call_type + "    " + a.Call_value)
	return true
}
