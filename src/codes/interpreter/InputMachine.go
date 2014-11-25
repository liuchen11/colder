package interpreter

/***************************************************************************************
**文件名：InputMachine.go
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
	"io"
	"os"
	"strconv"
)

type InputMachine struct {
	Ch         byte          //刚刚读入的字符
	Line       []byte        //当前行
	Length     int           //当前行的长度
	XPosition  int           //当前光标所在的位置
	YPosition  int           //当前光标所在的行数
	Endflag    bool          //结束标志
	Stringflag bool          //用于处理字符串
	Word       *Token        //当前的符号种类
	Keywords   []string      //关键字列表
	Reserved   []string      //保留字列表
	Operators  []string      //运算符列表
	Input      *bufio.Reader //输入文件流
}

//构造函数
func NewInputMachine(reader *bufio.Reader) *InputMachine {
	ret := new(InputMachine)
	ret.Input = reader
	ret.XPosition = 0
	ret.YPosition = 0
	ret.Keywords = make([]string, 0, 20)
	ret.Keywords = append(ret.Keywords, "bool", "break", "case", "default", "end", "false", "for", "int", "rbegin", "return", "string", "switch", "true", "void")
	ret.Reserved = make([]string, 0, 20)
	ret.Reserved = append(ret.Reserved, "get_cards", "get_cardsnum", "top_card", "find", "match", "deliver", "next_round", "compare", "play_card", "enable", "disable", "state_point", "state_card", "show_public", "show_private", "add_score")
	ret.Operators = make([]string, 0, 20)
	ret.Operators = append(ret.Operators, ">", "<", "=", "!", ">=", "<=", "==", "!=", "&&", "||", "+", "-", "*", "/", "%", "(", ")", "[", "]", "{", "}", ",", ";", "\"", ":")
	ret.Endflag = false
	ret.GetNextChar()
	return ret
}

//判断是否是空白符
func IsWhite(input byte) bool {
	if input == '\n' || input == '\t' || input == ' ' {
		return true
	} else {
		return false
	}
}

//向前读入一个字符，为了减少IO，每一读入一行
func (a *InputMachine) GetNextChar() (byte, error) {
	if a.Length == a.XPosition {
		for {
			a.Stringflag = false
			buffer, _, err := a.Input.ReadLine()
			a.Line = buffer
			if err == io.EOF {
				a.Endflag = true
				return 0, err
			} else if err != nil {
				fmt.Println("error")
				a.Endflag = true
				return 0, err
			}
			if len(buffer) != 0 {
				break
			}
		}
		a.Length = len(a.Line)
		a.XPosition = 0
		a.YPosition++
	}
	a.Ch = a.Line[a.XPosition]
	a.XPosition++
	//fmt.Println(a.Ch)
	return a.Ch, nil
}

//向前回退一格，如果当前光标所在位置在句首，则返回失败
func (a *InputMachine) BackSpace() bool {
	if a.XPosition == 1 {
		return false
	}
	a.XPosition--
	a.Ch = a.Line[a.XPosition]
	return true
}

//匹配
func (a *InputMachine) MatchNextToken() bool {
	for IsWhite(a.Ch) {
		_, err := a.GetNextChar()
		if err != nil {
			return false
		}
	}
	if a.Endflag == true {
		return false
	}
	var newToken bool
	if a.Stringflag == true && a.Word.Content == "\"" {
		newToken = a.MatchString()
	} else if (a.Ch >= 'a' && a.Ch <= 'z') || (a.Ch >= 'A' && a.Ch <= 'Z') || a.Ch == '_' {
		newToken = a.MatchCharacter()
	} else if a.Ch >= '0' && a.Ch <= '9' {
		newToken = a.MatchNumber()
	} else {
		newToken = a.MatchOperator()
	}
	if newToken == true {
		fmt.Println("Match a ", a.Word.Type, "->", a.Word.Content, "At ", strconv.FormatInt(int64(a.Word.Y+1), 10), " ", strconv.FormatInt(int64(a.Word.X), 10))
	}
	return newToken
}

//匹配字符串
func (a *InputMachine) MatchString() bool {
	var x int = a.XPosition
	var y int = a.YPosition
	var buffer []byte = make([]byte, 0, 10)
	for a.Ch != '"' && a.XPosition != 1 {
		buffer = append(buffer, a.Ch)
		_, err := a.GetNextChar()
		if err != nil {
			break
		}
	}
	a.Word = new(Token)
	var sbuffer string
	if len(buffer) == 0 {
		sbuffer = ""
	} else {
		sbuffer = string(buffer)
	}
	a.Word = new(Token)
	a.Word.Content = sbuffer
	a.Word.Type = "String"
	a.Word.X = x
	a.Word.Y = y
	return true
}

//匹配标识符
func (a *InputMachine) MatchCharacter() bool {
	var x int = a.XPosition
	var y int = a.YPosition
	var buffer []byte = make([]byte, 0, 10)
	for (len(buffer) == 0 || a.XPosition != 1) && ((a.Ch >= 'a' && a.Ch <= 'z') || (a.Ch >= 'A' && a.Ch <= 'Z') || (a.Ch >= '0' && a.Ch <= '9') || a.Ch == '_') {
		buffer = append(buffer, a.Ch)
		_, err := a.GetNextChar()
		if err != nil {
			break
		}
	}
	var sbuffer string = string(buffer)
	var iskeyword bool = false
	var isreservedword bool = false
	a.Word = new(Token)
	a.Word.Content = sbuffer
	var i int
	for i = 0; i < len(a.Keywords); i++ {
		if a.Keywords[i] == sbuffer {
			a.Word.Type = "Keyword"
			a.Word.Index = i
			iskeyword = true
			break
		}
	}
	for i = 0; i < len(a.Reserved); i++ {
		if a.Reserved[i] == sbuffer {
			a.Word.Type = "Reservedword"
			a.Word.Index = i
			isreservedword = true
			break
		}
	}
	if iskeyword == false && isreservedword == false {
		a.Word.Type = "Identifier"
	}
	a.Word.X = x
	a.Word.Y = y
	return true
}

//匹配数字。包括整型数字和浮点型数字
func (a *InputMachine) MatchNumber() bool {
	var x int = a.XPosition
	var y int = a.YPosition
	var value int = 0
	var buffer []byte = make([]byte, 0, 10)
	for a.Ch >= '0' && a.Ch <= '9' && (len(buffer) == 0 || a.XPosition != 1) {
		buffer = append(buffer, a.Ch)
		value = value*10 + int(a.Ch-'0')
		_, err := a.GetNextChar()
		if err != nil {
			break
		}
	}
	a.Word = new(Token)
	a.Word.Content = string(buffer)
	a.Word.Type = "Int"
	a.Word.IntValue = value
	a.Word.X = x
	a.Word.Y = y
	return true
}

//匹配符号，包括单字符符号和双字符符号
func (a *InputMachine) MatchOperator() bool {
	var x int = a.XPosition
	var y int = a.YPosition
	switch a.Ch {
	case '=':
		a.GetNextChar()
		if a.Ch == '=' && a.XPosition != 1 {
			a.Word = new(Token)
			a.Word.Content = "=="
			a.Word.Type = "Operator"
			a.Word.Index = 6
			a.GetNextChar()
		} else {
			a.Word = new(Token)
			a.Word.Content = "="
			a.Word.Type = "Operator"
			a.Word.Index = 2
		}
	case '>':
		a.GetNextChar()
		if a.Ch == '=' && a.XPosition != 1 {
			a.Word = new(Token)
			a.Word.Content = ">="
			a.Word.Type = "Operator"
			a.Word.Index = 4
			a.GetNextChar()
		} else {
			a.Word = new(Token)
			a.Word.Content = ">"
			a.Word.Type = "Operator"
			a.Word.Index = 0
		}
	case '<':
		a.GetNextChar()
		if a.Ch == '=' && a.XPosition != 1 {
			a.Word = new(Token)
			a.Word.Content = "<="
			a.Word.Type = "Operator"
			a.Word.Index = 5
			a.GetNextChar()
		} else {
			a.Word = new(Token)
			a.Word.Content = "<"
			a.Word.Type = "Operator"
			a.Word.Index = 1
		}
	case '!':
		a.GetNextChar()
		if a.Ch == '=' && a.XPosition != 1 {
			a.Word = new(Token)
			a.Word.Content = "!="
			a.Word.Type = "Operator"
			a.Word.Index = 7
			a.GetNextChar()
		} else {
			a.Word = new(Token)
			a.Word.Content = "!"
			a.Word.Type = "Operator"
			a.Word.Index = 3
		}
	case '&':
		a.GetNextChar()
		if a.Ch == '&' && a.XPosition != 1 {
			a.Word = new(Token)
			a.Word.Content = "&&"
			a.Word.Type = "Operator"
			a.Word.Index = 8
			a.GetNextChar()
		} else {
			fmt.Println("Error! At ", strconv.FormatInt(int64(a.Word.Y+1), 10), ":", strconv.FormatInt(int64(a.Word.X), 10))
			a.Endflag = true
			return false
		}
	case '|':
		a.GetNextChar()
		if a.Ch == '|' && a.XPosition != 1 {
			a.Word = new(Token)
			a.Word.Content = "||"
			a.Word.Type = "Operator"
			a.Word.Index = 9
			a.GetNextChar()
		} else {
			fmt.Println("Error! At ", strconv.FormatInt(int64(a.Word.Y+1), 10), ":", strconv.FormatInt(int64(a.Word.X), 10))
			a.Endflag = true
			return false
		}
	case '/':
		a.GetNextChar()
		if a.Ch == '/' && a.XPosition != 1 {
			a.XPosition = a.Length
			a.GetNextChar()
			return false
		} else {
			a.Word = new(Token)
			a.Word.Content = "/"
			a.Word.Type = "Operator"
			a.Word.Index = 13
		}
	default:
		var sbuffer string = string(a.Ch)
		var i int
		var isMatched bool = false
		for i = 10; i < len(a.Operators); i++ {
			if a.Operators[i] == sbuffer {
				if sbuffer == "\"" {
					a.Stringflag = !a.Stringflag
				}
				a.Word = new(Token)
				a.Word.Content = sbuffer
				a.Word.Type = "Operator"
				a.Word.Index = i
				a.GetNextChar()
				isMatched = true
				break
			}
		}
		if isMatched == false {
			fmt.Println("Error! At ", strconv.FormatInt(int64(a.Word.Y+1), 10), ":", strconv.FormatInt(int64(a.Word.X), 10))
			a.Endflag = true
			return false
		}
	}
	a.Word.X = x
	a.Word.Y = y
	return true
}

//与外部的接口函数，输入一个文件，输出一个Token数组
func Analyse(infile string) []Token {
	ret := make([]Token, 0, 10)
	fin, err := os.Open(infile)
	defer fin.Close()
	if err != nil {
		fmt.Println("Unable to open the file ", infile)
	}
	reader := bufio.NewReader(fin)
	machine := NewInputMachine(reader)
	for {
		ok := machine.MatchNextToken()
		if ok == true {
			ret = append(ret, *machine.Word)
		}
		if machine.Endflag == true {
			break
		}
	}
	return ret
}

