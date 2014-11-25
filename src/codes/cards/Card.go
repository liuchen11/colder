package cards

/***************************************************************************************
**文件名：Card.go
**包名称：cards
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


import (
	"fmt"
	"strings"
)

//表示卡牌的类型
type Card struct {
	Colour string
	Point  string
}

//转换为字符串
func (a *Card) ToString() string {
	return a.Colour + a.Point
}

//新建卡牌
func NewCard(str string) *Card {
	if len(str) < 2 {
		return nil
	}
	ret := new(Card)
	ret.Colour = str[0:1]
	ret.Point = str[1:2]
	return ret
}

//花色增加
func ColourIncrease(input string) string {
	switch input {
	case "J":
		return "J"
	case "H":
		return "D"
	case "D":
		return "S"
	case "S":
		return "C"
	case "C":
		return "H"
	default:
		return input
	}
}

//返回输入点数的下一个点数
func PointIncrease(input string) string {
	switch input {
	case "A":
		return "2"
	case "2":
		return "3"
	case "3":
		return "4"
	case "4":
		return "5"
	case "5":
		return "6"
	case "6":
		return "7"
	case "7":
		return "8"
	case "8":
		return "9"
	case "9":
		return "0"
	case "0":
		return "J"
	case "J":
		return "Q"
	case "Q":
		return "K"
	case "K":
		return "A"
	default:
		return input
	}
}

//将整型转化为花色
func Int2Colour(input int) string {
	input = input % 5
	switch input {
	case 0:
		return "J"
	case 1:
		return "H"
	case 2:
		return "D"
	case 3:
		return "S"
	case 4:
		return "C"
	default:
		return ""
	}
}

//将整型转化为点数
func Int2Point(input int) string {
	input = (input-1)%13 + 1
	switch input {
	case 1:
		return "A"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	case 10:
		return "0"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return ""
	}
}

//按照整型索引新建卡牌
//花色 0-王牌 1-红桃 2-方片 3-黑桃 4-草花
func NewCardByIndex(colour int, point int) *Card {
	var c string
	var p string
	switch colour {
	case 0:
		c = "J"
	case 1:
		c = "H"
	case 2:
		c = "D"
	case 3:
		c = "S"
	default:
		c = "C"
	}
	switch point {
	case 10:
		p = "0"
	case 11:
		p = "J"
	case 12:
		p = "Q"
	case 13:
		p = "K"
	case 1:
		if colour == 0 {
			p = "S"
		} else {
			p = "A"
		}
	case 2:
		if colour == 0 {
			p = "B"
		} else {
			p = "2"
		}
	default:
		p = string('0' + point)
	}
	ret := new(Card)
	ret.Colour = c
	ret.Point = p
	return ret
}

//获取卡牌点数索引
func (a *Card) GetPointIndex() int {
	switch a.Point {
	case "A":
		return 1
	case "0":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "B":
		return 2
	case "S":
		return 1
	default:
		return int(a.Point[0] - '0')
	}
}

//获取卡牌花色索引
func (a *Card) GetColourIndex() int {
	switch a.Colour {
	case "H":
		return 1
	case "D":
		return 2
	case "S":
		return 3
	case "C":
		return 4
	default:
		return 0
	}
}

//打印卡牌
func (a *Card) Print() {
	var c string
	var p string
	switch a.Colour {
	case "H":
		c = "红桃"
	case "D":
		c = "方片"
	case "S":
		c = "黑桃"
	case "C":
		c = "草花"
	default:
		if a.Point == "B" {
			fmt.Print("大王")
		} else {
			fmt.Print("小王")
		}
		return
	}
	switch a.Point {
	case "1":
		p = "A"
	case "0":
		p = "10"
	default:
		p = a.Point
	}
	fmt.Print(c, p)
	return
}

/*
将表示单个元素的正则表达式进行切分
返回值分别为表示花色的部分、表示点数的部分和该表达式是否合法
*/
func BreakRegEx(s string) (string, string, bool) {
	var pt string
	var cl string
	var split int
	s = strings.Replace(s, " ", "", -1)
	if len(s) <= 0 {
		return "", "", false
	}
	switch s[0] {
	case '[': //符合表达式表达花色，读取到]结束
		for i := 1; ; i++ {
			if i == len(s) {
				fmt.Println("syntax error in regular expression: ", s)
				return "", "", false
			}
			if s[i] == ']' {
				cl = s[1:i]
				split = i + 1
				break
			}
		}
	case '%': //用于引用全局变量的 如%colour%.
		for i := 1; ; i++ {
			if i == len(s) {
				fmt.Println("syntax error in regular expression: ", s)
				return "", "", false
			}
			if s[i] == '%' {
				if len(s) > i+1 && (s[i+1] == '#' || s[i+1] == '$') {
					i++
				}
				cl = s[0 : i+1]
				split = i + 1
				break
			}
		}
	default: //单个符号，需要考虑$ #等符号
		if len(s) > 1 && (s[1] == '#' || s[1] == '$') {
			split = 2
		} else {
			split = 1
		}
		cl = s[0:split]
	}
	if len(s) <= split {
		return cl, "", false
	}
	switch s[split] {
	case '[': //逻辑表达式表示点数
		for i := split + 1; ; i++ {
			if i == len(s) {
				fmt.Print("syntax error in regular expression: ", s)
				return cl, "", false
			}
			if s[i] == ']' {
				pt = s[split+1 : i]
				break
			}
		}
	case '%': //用于引用全局变量
		for i := split + 1; ; i++ {
			if i == len(s) {
				fmt.Print("syntax error in regular expression: ", s)
				return cl, "", false
			}
			if s[i] == '%' {
				if len(s) > i+1 && (s[i+1] == '#' || s[i+1] == '$') {
					i++
				}
				pt = s[split : i+1]
				break
			}
		}
	default: //单个变量或者常量
		if len(s) > split+1 && (s[split+1] == '#' || s[split+1] == '$') {
			pt = s[split : split+2]
		} else {
			pt = s[split : split+1]
		}
	}
	return cl, pt, true
}

//单张卡牌的正则表达式匹配，返回匹配与否
func (a *Card) Match(cl string, pt string) bool {
	return TestAndLogic(a.Point, pt) && TestAndLogic(a.Colour, cl)
}

//分析与关系
func TestAndLogic(input string, regex string) bool {
	regex_vec := strings.Split(regex, "&")
	for i := 0; i < len(regex_vec); i++ {
		if TestOrLogic(input, regex_vec[i]) == false {
			return false
		}
	}
	return true
}

//分析或关系
func TestOrLogic(input string, regex string) bool {
	regex_vec := strings.Split(regex, "|")
	for i := 0; i < len(regex_vec); i++ {
		if TestLogic(input, regex_vec[i]) == true {
			return true
		}
	}
	return false
}

//分析非关系以及单个纸牌的匹配
func TestLogic(input string, regex string) bool {
	var nec int = 0
	if regex[0] == '^' {
		nec = 1
	}
	if len(regex)-nec < 1 {
		fmt.Println("Syntax error in regular expression\n")
		return false
	}
	if regex[nec] == '.' || regex[nec] == input[0] {
		return nec == 0
	} else {
		return nec == 1
	}
}

//给一个模式串，返回左右符合要求的单张牌
func FindAllCards(mode string) []Card {
	fmt.Println("BEGIN")
	ret := make([]Card, 0, 10)
	colour, point, err := BreakRegEx(mode)
	fmt.Println(colour, point, err)
	if err == false {
		return ret
	}
	allcards := GenSetCards(1)
	for i := 0; i < len(allcards.Cards); i++ {
		result := allcards.Cards[i].Match(colour, point)
		if result == true {
			ret = append(ret, allcards.Cards[i])
		}
	}
	return ret
}

func Equals(c1 Card, c2 Card) bool {
	return c1.Point == c2.Point && c1.Colour == c2.Colour
}

