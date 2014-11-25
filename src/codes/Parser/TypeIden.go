package parser

/***************************************************************************************
**文件名：TypeIden.go
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

type TypeIden struct {
	typeiden_next *TypeIden
	typeiden_type string

	Typeiden_len     int
	Typeiden_type    string
	Typeiden_acttype string
}

func TypeIdenError() {
	fmt.Println("There is an error in typeiden...")
}
func TypeIdenRunError() {
	fmt.Println("RunTimeError:There is an errr in TypeIden...")
}

func GetTypeIden(tokenlist []Token) (*TypeIden, bool) {
	fmt.Print("TYPEIDEN:")
	showTokenlist(tokenlist)

	ret := new(TypeIden)
	size := len(tokenlist)
	if size == 1 {
		if tokenlist[0].Type == "Keyword" {
			ret.typeiden_type = "single"
			switch tokenlist[0].Content {
			case "int":
				ret.typeiden_type = "int"
			case "bool":
				ret.typeiden_type = "bool"
			case "string":
				ret.typeiden_type = "string"
			default:
				TypeIdenError()
				return nil, false
			}
			return ret, true
		} else {
			TypeIdenError()
			return nil, false
		}
	} else {
		if size >= 3 && tokenlist[size-1].Content == "]" && tokenlist[size-2].Content == "[" && tokenlist[size-1].Type == "Operator" && tokenlist[size-2].Type == "Operator" {
			tmp := tokenlist[0 : size-2]
			flag := false
			ret.typeiden_next, flag = GetTypeIden(tmp)
			//判断是否在下面出错，完成栈式报错
			if flag == false {
				TypeIdenError()
				return nil, false
			}
			ret.typeiden_type = "array"
			return ret, true
		}
	}
	return ret, true
}

func (a *TypeIden) Exe(na *NameTable, fu *FuncTable) bool {
	if a.typeiden_type == "array" {
		flag := a.typeiden_next.Exe(na, fu)
		//判断是否在下面出错，完成栈式报错
		if flag == false {
			TypeDefRunError()
			return false
		}
		a.Typeiden_type = a.typeiden_next.Typeiden_type
		a.Typeiden_len = a.typeiden_next.Typeiden_len + 1
		a.Typeiden_acttype = a.typeiden_next.Typeiden_acttype + "arr"
		return true
	} else {
		a.Typeiden_type = a.typeiden_type
		a.Typeiden_acttype = a.typeiden_type
		a.Typeiden_len = 0
		return true
	}
}
func (a *TypeIden) Show() {
	fmt.Print(a.Typeiden_type + " : ")
	fmt.Println(a.Typeiden_len)
}
