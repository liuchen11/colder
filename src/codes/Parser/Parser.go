package parser

/***************************************************************************************
**文件名：Parser.go
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

type Parser struct {
	content []Token
}

func NewParser(filename string) *Parser {
	ret := new(Parser)
	ret.content = Analyse(filename)
	return ret
}

func (ret *Parser) Show() {
	fmt.Println("==========================BEGIN TOKEN==========================")
	for k, v := range ret.content {
		fmt.Println("Token " + strconv.Itoa(k) + "  :  " + v.Type + "   " + v.Content)
	}
	fmt.Println("===========================END TOKEN===========================")
}

func (ret *Parser) Content() []Token {
	return ret.content
}

