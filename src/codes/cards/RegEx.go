package cards

/***************************************************************************************
**文件名：RegEx.go
**包名称：cards
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


import (
	"codes/tools"
	"fmt"
	"strconv"
	"strings"
)

//正则表达式类
type RegEx struct {
	Value string        //表达式
	vars  map[uint8]int //变量列表，用于表示重复次数
	parts []regex_range //组成成分
}

//测试
func (a *RegEx) GetParts() []regex_range {
	return a.parts
}

//用于测试，打印正则表达式的相关信息
func (a *RegEx) Print() {
	fmt.Println("-----------------------------")
	fmt.Println("value: ", a.Value)
	fmt.Println("vars:")
	for k, v := range a.vars {
		fmt.Println(k, " : ", v)
	}
	fmt.Println("parts:")
	for i := 0; i < len(a.parts); i++ {
		a.parts[i].Print()
	}
}

//用于表达正则表达式中的变量
type regex_var struct {
	value string //值
	mode  int    //类型 1表示花色 2表示点数 负数表示重复次数
	hasA  bool   //用于表述顺子，有没有K元素
}

//构造函数
func NewRegex_Var(v string, m int, A bool) *regex_var {
	ret := new(regex_var)
	ret.value = v
	ret.mode = m
	ret.hasA = A
	return ret
}

//用于测试
func (a *regex_var) Print() {
	fmt.Print("value: ", a.value, " mode: ", a.mode, " hasA: ", a.hasA, "\n")
}

//用于表示正则表达式中一个状态单元
//正则表达式最终转化为一个自动机
type regex_range struct {
	fix     bool     //出现次数是否确定
	times   int      //（至少）出现次数
	length  int      //包含的牌的张数
	content []string //牌的内容
	symbol  uint8    //对出现次数不是确定的牌组有效，出现的次数可能作为一个变量存储起来
}

//向状态单元中加入一个字符
func (a *regex_range) add(toadd string) {
	a.length++
	var s []string = make([]string, 0, 1)
	s = append(s, toadd)
	a.content = append(s, a.content...)
}

//测试，打印相关信息
func (a *regex_range) Print() {
	fmt.Print(a.fix, " ", a.times, " | ")
	for i := 0; i < a.length; i++ {
		fmt.Print(" ", a.content[i])
	}
	fmt.Printf("| symbol:%c", a.symbol)
	fmt.Print("\n")
}

//构造函数
func NewRegEx(init string) *RegEx {
	ret := new(RegEx)
	ret.Value = init
	ret.vars = make(map[uint8]int)
	ret.parts = make([]regex_range, 0, 3)
	return ret
}

/*
此步骤实现正则表达式的标准化
消除正则表达式中{}中的确定数字
把确定的数字罗列成枚举形式
*/
func (a *RegEx) Standardize() bool {
	stack := *new(tools.Stack) //用于处理左右括弧的匹配问题，存放内容为一对括弧的左侧位置
	stack.Clear()
	var std string = ""
	for i := 0; i < len(a.Value); i++ {
		switch a.Value[i] {
		case '\t':
			continue
		case '\n':
			continue
		case ' ':
			continue
		case ')':
			start := stack.Top().(int) //start表示右括弧对应的左括弧的位置，从而确定括弧中的内容
			stack.Pop()
			std = std + a.Value[i:i+1]
			i++
			var position int = i
			number := 0
			for state := 0; i < len(a.Value); i++ {
				if a.Value[i] == '\t' || a.Value[i] == '\n' || a.Value[i] == ' ' {
					continue
				} else if state == 0 {
					if a.Value[i] == '{' {
						state = 1
					} else {
						i--
						break
					}
				} else if state == 1 {
					if a.Value[i] >= '0' && a.Value[i] <= '9' {
						number = number*10 + int(a.Value[i]-'0')
					} else if a.Value[i] == '}' {
						recursion := NewRegEx(a.Value[start:position])
						recursion.Standardize() //递归，对括号内的内容进行标准化
						for j := 0; j < number-1; j++ {
							std = std + recursion.Value
						}
						break
					} else {
						for j := position; j <= i; j++ {
							if a.Value[j] != '\t' && a.Value[j] != '\n' && a.Value[j] != ' ' {
								std = std + a.Value[j:j+1]
							}
						}
						break
					}
				}
			}
		case '(':
			stack.Push(i) //匹配做左括弧，向栈中压入该左括弧的位置
			std = std + a.Value[i:i+1]
		default:
			std = std + a.Value[i:i+1]
		}
	}
	a.Value = std
	return true
}

//用于处理正则表达式中的最小单元，也就是表示单张纸牌
type unit struct {
	position int    //位置
	content  string //实际内容
}

/*
将正则表达式转化为若干个自动机单元
自动机单元的类型为regex_range,
每个单元包括内容、重复次数，重复次数是否确定等等信息
整个正则表达式匹配的过程视作一系列自动机状态的相互转化
*/
func (a *RegEx) Compile() bool {
	a.vars = make(map[uint8]int)
	var buf unit = *new(unit)
	var index tools.Stack    //用于左右括弧之间的匹配
	var item tools.Stack     //用于存放尚未划入状态单元的单张纸牌
	var new_unit bool = true //是否是新单元
	index.Clear()
	item.Clear()
	buf.content = ""
	for i := 0; i < len(a.Value); i++ {
		switch a.Value[i] {
		case '(':
			index.Push(i) //匹配左括弧
		case ')':
			start := index.Top().(int)
			index.Pop()
			if len(buf.content) > 1 {
				item.Push(buf)
				buf = *new(unit)
				buf.content = ""
				new_unit = true
			}
			if i+1 < len(a.Value) && a.Value[i+1] == '{' { //解析重复的次数
				i += 2
				range_now := *new(regex_range)
				for !item.Empty() && (item.Top().(unit)).position > start {
					var e string = (item.Top().(unit)).content
					range_now.add(e)
					item.Pop()
				}
				if a.Value[i] >= 'a' && a.Value[i] <= 'z' { //预定义次数
					_, ok := a.vars[a.Value[i]]
					if ok == false {
						fmt.Println("unrecognized symbol ", a.Value[i:i+1], " at ", i, " of ", a.Value)
						return false
					} else {
						range_now.fix = true
						range_now.times = -1 //重复次数为负数表示已经适应字母进行预定义
						range_now.symbol = a.Value[i]
						i++
					}
				} else {
					for number := 0; ; i++ {
						if a.Value[i] <= '9' && a.Value[i] >= '0' {
							number = number*10 + int(a.Value[i]-'0') //解析重复的次数
						} else {
							if a.Value[i] == '+' {
								range_now.fix = false
								range_now.times = number
								i++
								if a.Value[i] == ':' { //出现相应的变量定义，关注变量的必需做到不重复
									i++
									if a.Value[i] >= 'a' && a.Value[i] <= 'z' {
										_, ok := a.vars[a.Value[i]]
										if ok == true { //重定义
											fmt.Println("redefine the symbol ", a.Value[i:i+1], " at ", i, " of ", a.Value)
											return false
										} else {
											a.vars[a.Value[i]] = 0
											range_now.symbol = a.Value[i]
											i++
										}
									} else {
										fmt.Println("illegal symbol ", a.Value[i:i+1], " at ", i, " of ", a.Value)
										return false
									}
								} else {
									range_now.symbol = 0
								}
								break
							} else {
								fmt.Println("illegal symbol ", a.Value[i:i+1], " at ", i, " of ", a.Value)
								return false
							}
						}
					}
				}
				if a.Value[i] != '}' { //缺少右括弧
					fmt.Println("expected } at ", i, " of ", a.Value)
					return false
				}
				a.parts = append(a.parts, range_now)
			} else {
				if index.Empty() { //两边匹配的括号没有被外面更大的括号括住，否则的话，被外面更大括号括住的单元会被合并
					range_now := *new(regex_range)
					for !item.Empty() {
						var e string = (item.Top().(unit)).content
						range_now.add(e)
						item.Pop()
					}
					range_now.fix = true
					range_now.times = 1
					range_now.symbol = 0
					a.parts = append(a.parts, range_now)
				}
			}
		default:
			if new_unit == true {
				buf.position = i
				new_unit = false
			}
			buf.content = buf.content + a.Value[i:i+1]
		}
	}
	return true
}

/*
对自动机的单元进行匹配
第一个返回值为退出代号，第二个参数表示出现异常的情况，其值为false的时候表示出现错误
在没有出现错误的情况下，0表示匹配成功，-1表示匹配失败,1表示部分匹配，卡牌没有用完
出现错误时，第一个返回值表示出现的问题代号
*/
func MatchUnit(start int, cards []Card, state *regex_range, varlist *map[uint8]regex_var) (int, bool) {
	for i := 0; ; i++ {
		if i == state.length {
			if start == len(cards) {
				return 0, true //匹配成功
			} else {
				return 1, true //可以继续匹配
			}
		} else {
			if start == len(cards) {
				return -1, true //模式串没有用完，卡牌已经用完，匹配失败
			} else {
				parts := strings.Split(state.content[i], "||") //出现的卡牌中，可能出现逻辑或运算符，用||分开，匹配时，只要满则其中一一个进行了
				var hit bool = false
				for k := 0; k < len(parts); k++ {
					colour, point, ok := BreakRegEx(parts[k])
					if ok == false {
						return 1, false
					}
					for j := 'a'; j <= 'z'; j++ { //匹配卡牌中出现的变量
						var init string = strconv.QuoteRune(rune(j))[1:2]
						var inits string = init + "#"
						var dinit string = init + "$"
						_, exist := (*varlist)[uint8(j)]                                     //从参数列表中查找是否存在该变量
						if strings.Contains(colour, init) && strings.Contains(point, init) { //同一变量出现在花色和点数的位置
							fmt.Println("the symbol ", init, " appears in both colour period and point period in ", state.content[i])
							return 2, false
						}
						switch {
						case strings.Contains(colour, inits):
							if exist == false {
								colour = strings.Replace(colour, inits, ".", -1)
								(*varlist)[uint8(j)] = *NewRegex_Var(cards[start].Colour, 1, false)
							} else {
								if (*varlist)[uint8(j)].mode != 1 { //变量的类型不匹配，下同
									fmt.Println("mismatch the mode of the symbol ", init, ",", strconv.Itoa((*varlist)[uint8(j)].mode), " given,1 expected")
									return 3, false
								}
								var newvalue string = ColourIncrease((*varlist)[uint8(j)].value)
								(*varlist)[uint8(j)] = *NewRegex_Var(newvalue, 1, false)
								colour = strings.Replace(colour, inits, (*varlist)[uint8(j)].value, -1)
							}
						case strings.Contains(colour, dinit):
							colour = strings.Replace(colour, dinit, ".", -1)
							(*varlist)[uint8(j)] = *NewRegex_Var(cards[start].Colour, 1, false)
						case strings.Contains(colour, init):
							if exist == false {
								colour = strings.Replace(colour, init, ".", -1)
								(*varlist)[uint8(j)] = *NewRegex_Var(cards[start].Colour, 1, false)
							} else {
								if (*varlist)[uint8(j)].mode != 1 {
									fmt.Println("mismatch the mode of the symbol ", init, ",", strconv.Itoa((*varlist)[uint8(j)].mode), " given,1 expected")
									return 3, false
								}
								colour = strings.Replace(colour, init, (*varlist)[uint8(j)].value, -1)
							}
						case strings.Contains(point, inits):
							if exist == false {
								point = strings.Replace(point, inits, ".", -1)
								(*varlist)[uint8(j)] = *NewRegex_Var(cards[start].Point, 2, cards[start].Point == "A")
							} else {
								if (*varlist)[uint8(j)].mode != 2 {
									fmt.Println("mismatch the mode of the symbol ", init, ",", strconv.Itoa((*varlist)[uint8(j)].mode), " given,2 expected")
									return 3, false
								}
								newvalue := PointIncrease((*varlist)[uint8(j)].value)
								if newvalue == "A" {
									(*varlist)[uint8(j)] = *NewRegex_Var(newvalue, 2, true)
								} else {
									(*varlist)[uint8(j)] = *NewRegex_Var(newvalue, 2, (*varlist)[uint8(j)].hasA)
								}
								if (*varlist)[uint8(j)].hasA == true { //对于顺子进行特殊处理，因为K之后可以跟上A
									point = strings.Replace(point, inits, "K|"+(*varlist)[uint8(j)].value, -1)
								} else {
									point = strings.Replace(point, inits, (*varlist)[uint8(j)].value, -1)
								}
							}
						case strings.Contains(point, dinit):
							point = strings.Replace(point, dinit, ".", -1)
							(*varlist)[uint8(j)] = *NewRegex_Var(cards[start].Point, 2, cards[start].Point == "A")
						case strings.Contains(point, init):
							if exist == false {
								point = strings.Replace(point, init, ".", -1)
								(*varlist)[uint8(j)] = *NewRegex_Var(cards[start].Point, 2, cards[start].Point == "A")
							} else {
								if (*varlist)[uint8(j)].mode != 2 {
									fmt.Println("mismatch the mode of the symbol ", init, ",", strconv.Itoa((*varlist)[uint8(j)].mode), " given,2 expected")
									return 3, false
								}
								if (*varlist)[uint8(j)].hasA == true {
									point = strings.Replace(point, init, "K|"+(*varlist)[uint8(j)].value, -1)
								} else {
									point = strings.Replace(point, init, (*varlist)[uint8(j)].value, -1)
								}
							}
						}
					}
					match := cards[start].Match(colour, point)
					if match == true { //存在一个命中
						hit = true
						break
					}
				}
				if hit == false { //没有一个命中，匹配失败
					return -1, true
				}
			}
		}
		start++
	}
}

/*
正则表达式进行匹配
int返回参数表示退出代号,bool值表示匹配成功与否
*/
func (a *RegEx) Match(list CardList) (int, bool) {
	if a.Standardize() == false { //标准化
		fmt.Println("there are errors in the standardize process")
		return 1, false
	}
	if a.Compile() == false { //转化为一个单元一个单元的形式
		fmt.Println("there are errors in the compiling process")
		return 2, false
	}
	a.Print()
	var position []int = make([]int, 0, 5)                             //位置列表，每一个元素代表一种匹配方案
	var maps []map[uint8]regex_var = make([]map[uint8]regex_var, 0, 5) //符号表列表，每一个元素代表一个匹配方案
	cards := list.Cards
	position = append(position, 0)
	empty := make(map[uint8]regex_var)
	maps = append(maps, empty)
	//i 表示 正则表达式中 逻辑单元的数目
	for i := 0; i < len(a.parts); i++ {
		fmt.Println("i:", i)
		size := len(maps)
		if size == 0 {
			return 0, false
		}
		//j 表示可能匹配的方案总数
		for j := 0; j < size; {
			var repeated int = 0
			if a.parts[i].times < 0 { //重复的次数为变量的值
				sym := a.parts[i].symbol
				t, ok := maps[j][sym]
				if ok == false || t.mode != 3 {
					fmt.Printf("can not find the repeated times variable named %c\n", sym)
					fmt.Println("there are errors in the match process")
					return 3, false
				} else {
					repeated64, _ := strconv.ParseInt(t.value, 10, 32)
					repeated = int(repeated64)
				}
			} else {
				repeated = a.parts[i].times
			}
			if a.parts[i].fix == false && a.parts[i].symbol != 0 {
				newvar := *NewRegex_Var(strconv.FormatInt(int64(repeated), 10), 3, false)
				maps[j][uint8(a.parts[i].symbol)] = newvar
			}
			fmt.Println("repeated:", repeated)
			var drop bool = false
			for k := 0; k < repeated; k++ { //匹配至少重复的次数
				fmt.Println("OK")
				code, right := MatchUnit(position[j], cards, &a.parts[i], &maps[j])
				fmt.Println(code, right)
				if right == false {
					fmt.Println("there are errors in the match process")
					return 4, false
				} else {
					var end bool = false
					switch code {
					case 0:
						if i == len(a.parts)-1 && k == repeated-1 {
							return 0, true
						} else {
							position = append(position[:j], position[j+1:]...)
							maps = append(maps[:j], maps[j+1:]...)
							drop = true
							end = true
						}
					case 1:
						position[j] += a.parts[i].length
					default:
						position = append(position[:j], position[j+1:]...)
						maps = append(maps[:j], maps[j+1:]...)
						drop = true
						end = true
					}
					if len(position) == 0 {
						return 0, false
					}
					if end == true {
						fmt.Println("break")
						break
					}
				}
			}
			fmt.Println("OK")
			if drop == true {
				size--
			} else {
				j++
			}
		}
		fmt.Println("end")
		for j := 0; j < size; j++ {
			if a.parts[i].fix == false { //对于重复次数不确定的情况，超过重复次数下界之后可以继续进行匹配
				for {
					newmap := make(map[uint8]regex_var)
					for key, value := range maps[j] {
						newmap[key] = value
					}
					fmt.Println("OK1")
					code, right := MatchUnit(position[j], cards, &a.parts[i], &newmap)
					fmt.Println(code, right)
					var end bool = false
					if right == true {
						switch {
						case i == len(a.parts)-1 && code == 0:
							fmt.Println("1")
							return 0, true
						case code == 1:
							if a.parts[i].symbol != 0 {
								symbol_value, _ := strconv.ParseInt(newmap[a.parts[i].symbol].value, 10, 32)
								newmap[uint8(a.parts[i].symbol)] = *NewRegex_Var(strconv.FormatInt(symbol_value+1, 10), 3, false)
							}
							presentposition := position[j]
							position[j] += a.parts[i].length
							position = append(position, presentposition)
							presentmap := maps[j]
							maps[j] = newmap
							maps = append(maps, presentmap)
						default:
							fmt.Println("3")
							end = true
						}
						if end == true {
							break
						}
					} else {
						break
					}
				}
			}
		}
	}
	return 0, false
}

