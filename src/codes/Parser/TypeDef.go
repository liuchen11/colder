package parser

/***************************************************************************************
**文件名：TypeDef.go
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

type TypeDef struct {
	typedef_type string

	Typedef_type string
}

func TypeDefError() {
	fmt.Println("There is an error in TypeDef...")
}
func TypeDefRunError() {
	fmt.Println("There is an error in TypeDef...")
}

func GetTypeDef(tokenlist []Token) (*TypeDef, bool) {
	fmt.Print("TYPEDEF:")
	showTokenlist(tokenlist)
	ret := new(TypeDef)
	size := len(tokenlist)
	if size == 1 {
		if tokenlist[0].Type == "Keyword" {
			switch tokenlist[0].Content {
			case "int":
				ret.typedef_type = "int"
			case "bool":
				ret.typedef_type = "bool"
			case "string":
				ret.typedef_type = "string"
			default:
				TypeDefError()
				return nil, false
			}
			return ret, true
		} else {
			TypeDefError()
			return nil, false
		}
	} else {
		TypeDefError()
		return nil, false
	}
	return ret, true
}
func (a *TypeDef) Exe(na *NameTable, fu *FuncTable) bool {
	switch a.typedef_type {
	case "int":
		a.Typedef_type = "int"
		return true
	case "bool":
		a.Typedef_type = "bool"
		return true
	case "string":
		a.Typedef_type = "string"
		return true
	default:
		TypeDefRunError()
		return false
	}
	TypeDefRunError()
	return false
}

