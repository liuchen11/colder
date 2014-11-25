package interpreter

/***************************************************************************************
**文件名：Parser.go
**包名称：interpreter
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


import (
	"bufio"
	"fmt"
	"strconv"
)

type Parser struct {
	machine *InputMachine
	content []Token
}

func NewParser(reader *bufio.Reader) *Parser {
	ret := new(Parser)
	ret.machine = NewInputMachine(reader)
	ret.content = make([]Token, 0, 5)
	return ret
}

func (ret *Parser) SplitToken() {
	for {
		if ret.machine.Endflag == true {
			break
		}
		ret.machine.MatchNextToken()
		ret.content = append(ret.content, *(ret.machine.Word))
	}
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

