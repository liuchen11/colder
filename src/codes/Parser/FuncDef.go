package parser

/***************************************************************************************
**文件名：FuncDef.go
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

//类FuncDef
type FuncDef struct {
	funcdef_name     *Token
	funcdef_arg_type []*TypeIden
	funcdef_arg_name []*Token
	funcdef_content  *Stmt
	funcdef_type     *TypeIden
	Funcdef_name     string
	Funcdef_type     string
}

//报错
func FuncDefError() {
	fmt.Println("There is an error in FuncDef...")
}

//生成函数
func GetFuncDef(tokenlist []Token) (*FuncDef, bool) {
	fmt.Print("FUNCDEF:")
	showTokenlist(tokenlist)
	ret := new(FuncDef)
	size := len(tokenlist)
	i := 0
	//找到前面描述返回类型
	for i = 0; i < size; i++ {
		if tokenlist[i].Type == "Identifier" {
			break
		}
	}
	//如果没找到
	if i == size {
		FuncDefError()
		return nil, false
	}
	tmp := tokenlist[0:i]
	flag := false
	//解析类型
	ret.funcdef_type, flag = GetTypeIden(tmp)
	//判断是否在下面出错，完成栈式报错
	if flag == false {
		FuncDefError()
		return nil, false
	}
	//函数名赋值
	ret.funcdef_name = copy_token(&tokenlist[i])
	ret.Funcdef_name = ret.funcdef_name.Content
	i++
	if tokenlist[i].Content == "(" && tokenlist[i].Type == "Operator" {
		i++
		j := i
		//判断所有的参数列表
		for j = i; j < size; j++ {
			if tokenlist[j].Content == "," && tokenlist[j].Type == "Operator" {
				tmp := tokenlist[i : j-1]
				next_type, flag := GetTypeIden(tmp)
				//判断是否在下面出错，完成栈式报错
				if flag == false {
					FuncDefError()
					return nil, false
				}
				next_name := copy_token(&tokenlist[j-1])
				ret.funcdef_arg_type = append(ret.funcdef_arg_type, next_type)
				ret.funcdef_arg_name = append(ret.funcdef_arg_name, next_name)
				i = j + 1
				continue
			}
			if tokenlist[j].Content == ")" && tokenlist[j].Type == "Operator" {
				break
			}
		}
		//如果没有找到，报错
		if j == size {
			FuncDefError()
			return nil, false
		}
		//处理最后一个表达式
		if i != j {
			tmp := tokenlist[i : j-1]
			next_type, flag := GetTypeIden(tmp)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				FuncDefError()
				return nil, false
			}
			next_name := copy_token(&tokenlist[j-1])
			ret.funcdef_arg_type = append(ret.funcdef_arg_type, next_type)
			ret.funcdef_arg_name = append(ret.funcdef_arg_name, next_name)
		}
		//函数体的识别
		j++
		now := tokenlist[j:size]
		flag1 := false
		ret.funcdef_content, flag1 = GetStmt(now)
		if flag1 == false {
			FuncDefError()
			return nil, false
		}
		return ret, true
	} else {
		FuncDefError()
		return nil, false
	}
	return ret, true
}

//运行时候报错
func FuncDefRunError() {
	fmt.Println("RunTimeError:Error in FuncDef...")
}

//运行
func (a *FuncDef) Exe(na *NameTable, fu *FuncTable, in_type []string, in_value []string) bool {
	//添加到函数调用栈
	fu.AddFunc(a.Funcdef_name)

	flag1 := a.funcdef_type.Exe(na, fu)
	//栈式报错
	if flag1 == false {
		FuncDefRunError()
		return false
	}
	a.Funcdef_type = a.funcdef_type.Typeiden_acttype

	//Check
	size := len(a.funcdef_arg_type)
	size1 := len(a.funcdef_arg_name)
	//如果value和type数目不对等
	if size != size1 {
		FuncDefRunError()
		return false
	}
	sizea := len(in_type)
	sizeb := len(in_value)
	//如果输入的value和type个数不对应
	if sizea != sizeb {
		FuncDefRunError()
		return false
	}
	//如果参数和实际调用不同意
	if size != sizea {
		FuncDefRunError()
		return false
	}
	//复制符号表
	now := CopyNameTable(na)
	now.Show()
	//Add Arg
	//新建符号表
	for i := 0; i < size; i++ {
		type1 := in_type[i]
		flag := a.funcdef_arg_type[i].Exe(now, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			FuncDefRunError()
			return false
		}
		type2 := a.funcdef_arg_type[i].Typeiden_acttype
		if type1 != type2 {
			FuncDefRunError()
			return false
		}
		//函数名称
		v_name := a.funcdef_arg_name[i].Content
		//对于类型进行判断
		switch type1 {
		case "int":
			//返回整形
			now.AddVariable("int", v_name)
			v_value, _ := strconv.ParseInt(in_value[i], 10, 32)
			now.SetInt(v_name, int(v_value))
		case "bool":
			//返回为bool
			now.AddVariable("bool", v_name)
			v_value := false
			if in_value[i] == "true" {
				v_value = true
			}
			now.SetBool(v_name, v_value)
		case "string":
			//返回为string
			now.AddVariable("string", v_name)
			v_value := in_value[i]
			now.SetString(v_name, v_value)
		case "intarr":
			//对于整形
			//使用函数分隔开数组内容
			tmp := strings.Split(in_value[i], ",")
			//fmt.Println("==============")
			//fmt.Println(v_name)
			//fmt.Println("==============")
			arrsize := len(tmp)
			now.AddArray("intarr", v_name, arrsize)
			for i := 0; i < arrsize; i++ {
				value_int, _ := strconv.ParseInt(tmp[i], 10, 32)
				now.SetIntArr(v_name, i, int(value_int))
			}
		case "boolarr":
			//对于bool数组
			tmp := strings.Split(in_value[i], ",")
			arrsize := len(tmp)
			now.AddArray("boolarr", v_name, arrsize)
			for i := 0; i < arrsize; i++ {
				value := false
				if tmp[i] == "true" {
					value = true
				}
				//添加到符号表
				now.SetBoolArr(v_name, i, value)
			}
		case "stringarr":
			//对于string数组
			tmp := strings.Split(in_value[i], ",")
			arrsize := len(tmp)
			//新建一个变量
			now.AddArray("stringarr", v_name, arrsize)
			for i := 0; i < arrsize; i++ {
				//添加到符号表
				now.SetStringArr(v_name, i, tmp[i])
			}
		default:
			FactorRunError()
			return false
		}
	}
	//调试：显示复制后的NameTable
	//now.Show()
	//执行函数体
	flag := a.funcdef_content.Exe(now, fu)
	//判断是否在下面出错，完成栈式报错
	if flag == false {
		FuncDefRunError()
		return false
	}
	//now.Show()
	return true
}
