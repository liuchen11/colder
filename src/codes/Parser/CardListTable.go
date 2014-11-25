package parser

import (
	"codes/cards"
	"codes/tools"
	"fmt"
	"strconv"
	"strings"
)

//规则中的case语句
type condition_result struct {
	condition string
	result    int
}

func newcondition_result(c string, r int) *condition_result {
	ret := new(condition_result)
	ret.condition = c
	ret.result = r
	return ret
}

//每一个字符串匹配串
type comparelistline struct {
	name          string             //名称
	mode          cards.RegEx        //模式串
	compare_list  string             //比较列表
	post_prior    []string           //出牌上游类
	compare_prior []string           //比较上游类
	post_rule     []condition_result //出牌规则
	compare_rule  []condition_result //比较规则
}

//comparelistline的构造函数
func newcomparelistline(n string, m cards.RegEx, c_l string, p_p []string, c_p []string, p_r []condition_result, c_r []condition_result) *comparelistline {
	ret := new(comparelistline)
	ret.name = n
	ret.mode = m
	ret.compare_list = c_l
	ret.post_prior = p_p
	ret.compare_prior = c_p
	ret.post_rule = p_r
	ret.compare_rule = c_r
	return ret
}

//牌组比较列表
type CardListTable struct {
	Table CardTable
	Rule  []comparelistline
}

//CardListTable构造函数
func NewCardListTable() *CardListTable {
	ret := new(CardListTable)
	ret.Rule = make([]comparelistline, 0, 3)
	ret.Table = *NewCardTable()
	return ret
}

func (a *CardListTable) AddList(line string) bool { return a.Table.AddList(line) }

//加入新的模式串
func (a *CardListTable) AddMode(line string) bool {
	parts := strings.Split(line, ",")
	//parts = tools.WipeOutBlank(parts)
	if len(parts) < 7 {
		return false
	}
	name := parts[0]
	mode := *cards.NewRegEx(parts[1])
	compare_list := parts[2]
	post_prior := strings.Split(parts[3], "|")
	post_prior = tools.WipeOutBlank(post_prior)
	compare_prior := strings.Split(parts[4], "|")
	compare_prior = tools.WipeOutBlank(compare_prior)
	parts[5] = strings.Replace(parts[5], "}", "", -1)
	parts[5] = strings.Replace(parts[5], "{", "", -1)
	parts[6] = strings.Replace(parts[6], "}", "", -1)
	parts[6] = strings.Replace(parts[6], "{", "", -1)
	p_rules := strings.Split(parts[5], ";")
	c_rules := strings.Split(parts[6], ";")
	post_rule := make([]condition_result, 0, 3)
	compare_rule := make([]condition_result, 0, 3)
	for i := 0; i < len(p_rules); i++ {
		reason_result := strings.Split(p_rules[i], "=>")
		if len(reason_result) >= 2 {
			r, ok := strconv.ParseInt(reason_result[1], 10, 32)
			if ok == nil {
				toadd := *newcondition_result(reason_result[0], int(r))
				post_rule = append(post_rule, toadd)
			}
		}
	}
	for i := 0; i < len(c_rules); i++ {
		reason_result := strings.Split(c_rules[i], "=>")
		if len(reason_result) >= 2 {
			r, ok := strconv.ParseInt(reason_result[1], 10, 32)
			if ok == nil {
				toadd := *newcondition_result(reason_result[0], int(r))
				compare_rule = append(compare_rule, toadd)
			}
		}
	}
	toadd := newcomparelistline(name, mode, compare_list, post_prior, compare_prior, post_rule, compare_rule)
	fmt.Println(toadd.name, toadd.mode.Value, toadd.compare_list, toadd.post_prior, toadd.compare_prior)
	for i := 0; i < len(toadd.post_rule); i++ {
		fmt.Println(toadd.post_rule[i].condition, "=>", toadd.post_rule[i].result)
	}
	fmt.Println("---")
	for i := 0; i < len(toadd.compare_rule); i++ {
		fmt.Println(toadd.compare_rule[i].condition, "=>", toadd.compare_rule[i].result)
	}
	a.Rule = append(a.Rule, *toadd)
	return true
}

//对语句串进行切分
func SplitExpr(input string) []string {
	output := make([]string, 0, 3)
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '\n':
		case '\t':
		case ' ':
		case '>':
			if i+1 < len(input) && input[i+1] == '=' {
				output = append(output, input[i:i+2])
				i++
			} else {
				output = append(output, input[i:i+1])
			}
		case '!':
			if i+1 < len(input) && input[i+1] == '=' {
				output = append(output, input[i:i+2])
				i++
			} else {
				output = append(output, input[i:i+1])
			}
		case '<':
			if i+1 < len(input) && input[i+1] == '=' {
				output = append(output, input[i:i+2])
				i++
			} else {
				output = append(output, input[i:i+1])
			}
		case '=':
			if i+1 < len(input) && input[i+1] == '=' {
				output = append(output, input[i:i+2])
				i++
			} else {
				output = append(output, input[i:i+1])
			}
		case '(':
			output = append(output, input[i:i+1])
		case ')':
			output = append(output, input[i:i+1])
		case '.':
		case '$':
			output = append(output, input[i:i+1])
			i++
			start := i
			for input[i] >= '0' && input[i] <= '9' {
				i++
			}
			output = append(output, input[start:i])
			i--
		default:
			start := i
			for (input[i] >= 'a' && input[i] <= 'z') || (input[i] >= 'A' && input[i] <= 'Z') || (input[i] >= '0' && input[i] <= '9') {
				i++
			}
			output = append(output, input[start:i])
			i--
		}
	}
	return output
}

//判断语句条件是否成立
func (a *CardListTable) Justify(list1 cards.CardList, list2 cards.CardList, expr string, name string, nametable NameTable) (bool, bool) {
	if expr == "" || expr == "." {
		return true, true
	}
	parts := SplitExpr(expr)
	fmt.Println(parts)
	var rightlevel int = 0 //0表示牌组层面的 1表示单牌层面的 2表示花色点数层面的
	var leftlevel int = 0
	var presentlist int = 0 //当前考虑的牌组
	var size1 int           //牌组的大小
	var size2 int
	var card1 *cards.Card //卡牌
	var card2 *cards.Card
	var property1 string //属性
	var property2 string
	var operator int = -1 //操作符
	for i := 0; i < len(parts); i++ {
		fmt.Println(i)
		switch parts[i] {
		case "":
			fmt.Println("syntax error!")
			return false, false
		case "$":
			if i+1 >= len(parts) {
				fmt.Println("syntax error!")
			} else {
				i++
				number, ok := strconv.ParseInt(parts[i], 10, 32)
				if ok != nil {
					fmt.Println("syntax error!")
					return false, false
				} else {
					switch number {
					case 1:
						presentlist = 1
					case 2:
						presentlist = 2
					default:
						fmt.Println("cross the boundary!")
						return false, false
					}
				}
			}
		case "(":
			if rightlevel == 0 {
				rightlevel = 1
				if i+2 >= len(parts) {
					fmt.Println("syntax error!")
				} else {
					i++
					number, ok := strconv.ParseInt(parts[i], 10, 32)
					i++
					switch {
					case ok != nil:
						fallthrough
					case parts[i] != ")":
						fmt.Println("syntax error!")
						return false, false
					case int(number) <= len(list1.Cards) && presentlist == 1:
						card1 = &list1.Cards[number-1]
					case int(number) <= len(list2.Cards) && presentlist == 2:
						card2 = &list2.Cards[number-1]
					default:
						fmt.Println("syntax error!")
						return false, false
					}

				}
			} else {
				fmt.Println("syntax error!")
				return false, false
			}
		case "size":
			switch presentlist {
			case 1:
				size1 = len(list1.Cards)
			case 2:
				size2 = len(list1.Cards)
			default:
				fmt.Println("syntax error!")
				return false, false
			}
		case "colour":
			if rightlevel != 1 {
				rightlevel = 2
				switch presentlist {
				case 1:
					if card1 != nil {
						property1 = card1.Colour
					} else {
						fmt.Println("syntax error!")
						return false, false
					}
				case 2:
					if card2 != nil {
						property2 = card2.Colour
					} else {
						fmt.Println("syntax error!")
						return false, false
					}
				default:
					fmt.Println("syntax error!")
					return false, false
				}
			} else {
				fmt.Println("syntax error!")
				return false, false
			}
		case "point":
			if rightlevel != 1 {
				rightlevel = 2
				switch presentlist {
				case 1:
					if card1 != nil {
						property1 = card1.Point
					} else {
						fmt.Println("syntax error!")
						return false, false
					}
				case 2:
					if card2 != nil {
						property2 = card2.Point
					} else {
						fmt.Println("syntax error!")
						return false, false
					}
				default:
					fmt.Println("syntax error!")
					return false, false
				}
			} else {
				fmt.Println("syntax error!")
				return false, false
			}
		case "==":
			operator = 0
			leftlevel = rightlevel
			rightlevel = 0
			presentlist = 0
		case "!=":
			operator = 1
			leftlevel = rightlevel
			rightlevel = 0
			presentlist = 0
		case ">":
			operator = 2
			leftlevel = rightlevel
			rightlevel = 0
			presentlist = 0
		case ">=":
			operator = 3
			leftlevel = rightlevel
			rightlevel = 0
			presentlist = 0
		case "<":
			operator = 4
			leftlevel = rightlevel
			rightlevel = 0
			presentlist = 0
		case "<=":
			operator = 5
			leftlevel = rightlevel
			rightlevel = 0
			presentlist = 0
		default:
			fmt.Println("syntax error!")
			return false, false
		}
	}
	if leftlevel == rightlevel {
		switch leftlevel {
		case 0:
			switch operator {
			case 0:
				if size1 == size2 {
					return true, true
				} else {
					return false, true
				}
			case 1:
				if size1 != size2 {
					return true, true
				} else {
					return false, true
				}
			case 2:
				if size1 > size2 {
					return true, true
				} else {
					return false, true
				}
			case 3:
				if size1 >= size2 {
					return true, true
				} else {
					return false, true
				}
			case 4:
				if size1 < size2 {
					return true, true
				} else {
					return false, true
				}
			case 5:
				if size1 <= size2 {
					return true, true
				} else {
					return false, true
				}
			default:
				fmt.Println("unsupported operator!")
				return false, false
			}
		case 1:
			if card1 == nil || card2 == nil {
				fmt.Println("syntax error!")
				return false, false
			} else {
				result := a.Table.Compare(*card1, *card2, name, nametable)
				switch operator {
				case 0:
					if result == 0 {
						return true, true
					} else {
						return false, true
					}
				case 1:
					if result != 0 {
						return true, true
					} else {
						return false, true
					}
				case 2:
					if result == 1 {
						return true, true
					} else {
						return false, true
					}
				case 3:
					if result == 1 || result == 0 {
						return true, true
					} else {
						return false, true
					}
				case 4:
					if result == -1 {
						return true, true
					} else {
						return false, true
					}
				case 5:
					if result == 0 || result == -1 {
						return true, true
					} else {
						return false, true
					}
				default:
					fmt.Println("unsupported operator!")
					return false, false
				}
			}
		case 2:
			if property1 == "" || property2 == "" {
				fmt.Println("syntax error!")
				return false, false
			} else {
				switch operator {
				case 0:
					if property1 == property2 {
						return true, true
					} else {
						return false, true
					}
				case 1:
					if property1 != property2 {
						return true, true
					} else {
						return false, true
					}
				default:
					fmt.Println("unsupported operator!")
					return false, false
				}
			}
		default:
			fmt.Println("syntax error!")
			return false, false
		}
	} else {
		fmt.Println("the type of left is not the same of the right")
		return false, false
	}
}

//在出牌的情况下比较牌组的大小
func (a *CardListTable) Post(list1 cards.CardList, list2 cards.CardList, nametable NameTable) (int, bool) {
	var index1 int = -1
	var index2 int = -1
	for i := 0; i < len(a.Rule); i++ {
		//list1.Print()
		//fmt.Println(a.Rule[i].mode.Value)
		code, ok := list1.Match(a.Rule[i].mode)
		if code == 0 && ok == true {
			index1 = i
			break
		}
	}
	for i := 0; i < len(a.Rule); i++ {
		code, ok := list2.Match(a.Rule[i].mode)
		if code == 0 && ok == true {
			index2 = i
			break
		}
	}
	if index1 == -1 || index2 == -1 {
		return 2, true
	}
	fmt.Println(index1, index2)
	if index1 == index2 {
		for j := 0; j < len(a.Rule[index1].post_rule); j++ {
			fmt.Println(a.Rule[index1].post_rule[j].condition)
			result, normal := a.Justify(list1, list2, a.Rule[index1].post_rule[j].condition, a.Rule[index1].compare_list, nametable)
			if result == true && normal == true {
				return a.Rule[index1].post_rule[j].result, true
			}
		}
		return 2, true
	} else {
		name1 := a.Rule[index1].name
		name2 := a.Rule[index2].name
		prior1 := a.Rule[index1].post_prior
		prior2 := a.Rule[index2].post_prior
		for j := 0; j < len(prior1); j++ {
			if name2 == prior1[j] {
				return -1, true
			}
		}
		for j := 0; j < len(prior2); j++ {
			if name1 == prior2[j] {
				return 1, true
			}
		}
		return 2, true
	}
}

//比较两个牌组的大小
func (a *CardListTable) Compare(list1 cards.CardList, list2 cards.CardList, nametable NameTable) (int, bool) {
	var index1 int = -1
	var index2 int = -1
	for i := 0; i < len(a.Rule); i++ {
		code, ok := list1.Match(a.Rule[i].mode)
		fmt.Println(code, "----", ok)
		if code == 0 && ok == true {
			index1 = i
			break
		}
	}
	for i := 0; i < len(a.Rule); i++ {
		code, ok := list2.Match(a.Rule[i].mode)
		fmt.Println(code, "----", ok)
		if code == 0 && ok == true {
			index2 = i
			break
		}
	}
	if index1 == -1 || index2 == -1 {
		return 2, true
	}
	if index1 == index2 {
		for j := 0; j < len(a.Rule[index1].compare_rule); j++ {
			result, normal := a.Justify(list1, list2, a.Rule[index1].compare_rule[j].condition, a.Rule[index1].compare_list, nametable)
			if result == true && normal == true {
				return a.Rule[index1].post_rule[j].result, true
			}
		}
		return 2, true
	} else {
		name1 := a.Rule[index1].name
		name2 := a.Rule[index2].name
		prior1 := a.Rule[index1].compare_prior
		prior2 := a.Rule[index2].compare_prior
		for j := 0; j < len(prior1); j++ {
			if name2 == prior1[j] {
				return -1, true
			}
		}
		for j := 0; j < len(prior2); j++ {
			if name1 == prior2[j] {
				return 1, true
			}
		}
		return 2, true
	}
}
