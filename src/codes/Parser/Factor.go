package parser

/***************************************************************************************
**文件名：Factor.go
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

//类Factor
type Factor struct {
	factor_type  string //Factor为哪种Factor
	factor_call  *Call  //调用型的Factor
	factor_token *Token //单个变量或者是数组的名字
	factor_expr  *Expr  //表达式

	Factor_value_type   string //Factor的类型
	Factor_value_single string //Factor的值
}

//Factor 报错
func FactorError() {
	fmt.Println("Ther is an error in Factor....")
}

//GetFactor:生成一个Factor的指针

func GetFactor(tokenlist []Token) (*Factor, bool) {
	fmt.Print("FACTOR : ")
	showTokenlist(tokenlist)
	ret := new(Factor)
	size := len(tokenlist)
	//单个变量的factor
	if size == 1 {
		now := tokenlist[0]
		switch now.Type {
		case "Int":
			//整形的常量
			ret.factor_type = "intconstant"
			ret.factor_token = copy_token(&now)
			//fmt.Println("Int : " + now.Content)
			return ret, true
		case "Keyword":
			//布尔型常量
			if now.Content == "true" {
				ret.factor_type = "boolconstant"
				ret.factor_token = copy_token(&now)
				//	fmt.Println("True")
				return ret, true
			} else {
				if now.Content == "false" {
					ret.factor_type = "boolconstant"
					ret.factor_token = copy_token(&now)
					//		fmt.Println("False")
					return ret, true
				} else {
					FactorError()
					return nil, false
				}
			}
		case "Identifier":
			//一个变量：对应一个标示符
			ret.factor_token = copy_token(&now)
			ret.factor_type = "variable"
			//	fmt.Println("Variable")
			return ret, true

		default:
			FactorError()
			return nil, false
		}
	} else {
		//如果只有""，遇到一个字符串常量
		if size == 2 && tokenlist[0].Content == "\"" && tokenlist[1].Content == "\"" && tokenlist[0].Type == "Operator" && tokenlist[1].Type == "Operator" {
			ret.factor_type = "nonstringconstant"
			return ret, true
		} else {
			//如果遇到了一个中间不定
			//但是两侧都是引号的Token认为是字符串常量
			if size == 3 && tokenlist[0].Content == "\"" && tokenlist[0].Type == "Operator" && tokenlist[2].Type == "Operator" && tokenlist[2].Content == "\"" {
				ret.factor_token = copy_token(&(tokenlist[1]))
				ret.factor_type = "stringconstant"
				fmt.Println("Constant String")
				return ret, true
			} else {
				//遇到开头是标示符号或者是保留字的认为是调用类型的Factor
				if size > 2 && (tokenlist[0].Type == "Identifier" || tokenlist[0].Type == "Reservedword") && tokenlist[1].Type == "Operator" && tokenlist[1].Content == "(" && tokenlist[size-1].Content == ")" && tokenlist[size-1].Type == "Operator" {
					flag := false
					ret.factor_call, flag = GetCall(tokenlist)
					//判断是否在下面出错，完成栈式报错
					if flag == false {
						FactorError()
						return nil, false
					}
					ret.factor_type = "call"
					//fmt.Println("Caller")
					return ret, true
				} else {
					//同上，不过如果遇到的是中括号，认为是访问下标的Factor
					if size > 2 && tokenlist[0].Type == "Identifier" && tokenlist[1].Type == "Operator" && tokenlist[1].Content == "[" && tokenlist[size-1].Content == "]" && tokenlist[size-1].Type == "Operator" {
						tokenlist_new := tokenlist[2 : size-1]
						flag := false
						ret.factor_expr, flag = GetExpr(tokenlist_new)
						//判断是否在下面出错，完成栈式报错
						if flag == false {
							FactorError()
							return nil, false
						}
						ret.factor_token = copy_token(&tokenlist[0])
						ret.factor_type = "array"
						//fmt.Println("Array")
						return ret, true
					} else {
						//表达式类型的Factor
						if size >= 2 && tokenlist[0].Content == "(" && tokenlist[0].Type == "Operator" && tokenlist[size-1].Type == "Operator" && tokenlist[size-1].Content == ")" {
							flag := false
							ret.factor_type = "expr"
							ret.factor_expr, flag = GetExpr(tokenlist[1 : size-1])
							//判断是否在下面出错，完成栈式报错
							if flag == false {
								FactorError()
								return nil, false
							}
							return ret, true
						}
					}
				}
			}

		}
	}
	return ret, true
}

//调试使用：打印Factor的情况
func (a *Factor) Show(level int) {
	PrintIndent(level)
	fmt.Print("Factor:")
	switch a.factor_type {
	case "intconstant": //整型常量
		fmt.Println("Int : " + a.factor_token.Content)
	case "boolconstant": //bool常量
		fmt.Println("Bool : " + a.factor_token.Content)
	case "stringconstant": //String常量
		fmt.Println("String : " + a.factor_token.Content)
	case "nonstringconstant": //非空常量
		fmt.Println("EmptyString")
	case "variable": //变量
		fmt.Println("Variable : " + a.factor_token.Content)
	case "array": //数组
		fmt.Println("Array : " + a.factor_token.Content) //数组
		a.factor_expr.Show(level + 1)
	case "call":
		a.factor_expr.Show(level + 1)
	}
}

//Factor运行时报错
func FactorRunError() {
	fmt.Println("RunTimeError:There is an error in Factor...")
}

//Factor执行
func (a *Factor) Exe(na *NameTable, fu *FuncTable) bool {
	switch a.factor_type {
	case "intconstant":
		//整型常量
		a.Factor_value_type = "int"
		a.Factor_value_single = a.factor_token.Content
		return true
	case "boolconstant":
		//bool常量
		a.Factor_value_type = "bool"
		a.Factor_value_single = a.factor_token.Content
		return true
	case "stringconstant":
		//字符串常量
		a.Factor_value_type = "string"
		a.Factor_value_single = a.factor_token.Content
		return true
	case "variable":
		//变量类型
		type_f := na.GetType(a.factor_token.Content)
		//根据这个变量在的不同类型
		switch type_f {
		case "int":
			//对于整形变量
			a.Factor_value_type = "int"
			a.Factor_value_single = strconv.FormatInt(int64(na.GetInt(a.factor_token.Content)), 10)
			return true
		case "bool":
			//对于Bool变脸
			a.Factor_value_type = "bool"
			a.Factor_value_single = strconv.FormatBool(na.GetBool(a.factor_token.Content))
			return true
		case "string":
			//对于string类型变量
			a.Factor_value_type = "string"
			a.Factor_value_single = na.GetString(a.factor_token.Content)
			return true
		case "intarr":
			//对于整形数组
			//输出的结果应该是整个数组的值，并用，进行连接
			//需要提前求出长度
			a.Factor_value_type = "intarr"
			v_name := a.factor_token.Content
			size := na.GetIntlen(a.factor_token.Content)
			if size == 0 {
				a.Factor_value_single = ""
				return true
			}
			tmp := ""
			for i := 0; i < size; i++ {
				test, _ := na.GetIntArr(v_name, i)
				value := strconv.FormatInt(int64(test), 10)
				tmp = tmp + value + ","
			}
			//这里需要注意，应该忽略掉最后一个,
			tmp = tmp[0 : len(tmp)-1]
			a.Factor_value_single = tmp
			return true
		case "stringarr":
			//对于string数组
			//输出的结果应该是整个数组的值，并用，进行连接
			//需要提前求出长度
			a.Factor_value_type = "stringarr"
			v_name := a.factor_token.Content
			size := na.GetStringlen(a.factor_token.Content)
			if size == 0 {
				a.Factor_value_single = ""
				return true
			}
			tmp := ""
			for i := 0; i < size; i++ {
				value, _ := na.GetStringArr(v_name, i)
				tmp = tmp + value + ","
			}
			//这里需要注意，应该忽略掉最后一个,
			tmp = tmp[0 : len(tmp)-1]
			a.Factor_value_single = tmp
			return true
		case "boolarr":
			//对于bool数组
			//输出的结果应该是整个数组的值，并用，进行连接
			//需要提前求出长度
			a.Factor_value_type = "boolarr"
			v_name := a.factor_token.Content
			size := na.GetBoollen(a.factor_token.Content)
			if size == 0 {
				a.Factor_value_single = ""
				return true
			}
			tmp := ""
			for i := 0; i < size; i++ {
				test, _ := na.GetBoolArr(v_name, i)
				value := "true"
				if test == false {
					value = "false"
				}
				tmp = tmp + value + ","
			}
			//这里需要注意，应该忽略掉最后一个,
			tmp = tmp[0 : len(tmp)-1]
			a.Factor_value_single = tmp
			return true
		default:
			fmt.Println("UnSupported Array TYPE!!!!!")
			FactorRunError()
			return false
		}
		return true
	case "array":
		//数组调用类型
		array_name := a.factor_token.Content
		flag := a.factor_expr.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			FactorRunError()
			return false
		}
		must_int := a.factor_expr.Expr_value_type
		//数组下标的
		if must_int != "int" {
			FactorRunError()
			return false
		}
		array_type := na.GetType(array_name)
		label, _ := strconv.ParseInt(a.factor_expr.Expr_value, 10, 64)
		//调试使用
		//fmt.Println("BOMB~")
		//na.Show()
		//fmt.Println(label)
		//fmt.Println("bomb~")
		//对于不同类型的数组进行分类讨论
		switch array_type {
		case "intarr":
			//整形数组直接返回整形
			a.Factor_value_type = "int"
			tmp, _ := na.GetIntArr(array_name, int(label))
			value_int := strconv.FormatInt(int64(tmp), 10)
			a.Factor_value_single = value_int
			fmt.Println(a.Factor_value_single)
		case "boolarr":
			//Bool数组
			a.Factor_value_type = "bool"
			tmp, _ := na.GetBoolArr(array_name, int(label))
			value_bool := "false"
			if tmp == true {
				value_bool = "true"
			}
			a.Factor_value_single = value_bool
		case "stringarr":
			//string数组
			a.Factor_value_type = "string"
			tmp, _ := na.GetStringArr(array_name, int(label))
			value_string := tmp
			a.Factor_value_single = value_string
		default:
			FactorRunError()
			return false
		}
		return true
	case "call":
		//调用型Factor
		flag := a.factor_call.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			FactorRunError()
			return false
		}
		a.Factor_value_single = a.factor_call.Call_value
		a.Factor_value_type = a.factor_call.Call_type
		return true
	case "expr":
		//用括号括起来的表达式
		flag := a.factor_expr.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			FactorRunError()
			return false
		}
		a.Factor_value_type = a.factor_expr.Expr_value_type
		a.Factor_value_single = a.factor_expr.Expr_value
		return true
	default:
		fmt.Println("Exe Error in Factor.....")
		return false
	}
	return false
}
