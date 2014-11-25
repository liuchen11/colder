package parser

import (
	"codes/cards"
	"codes/tools"
	"fmt"
	"strings"
)

//卡牌优先级列表中的一行
type compareline struct {
	name     string   //列表名
	card1    string   //卡牌1
	card2    string   //卡牌2
	cardlist []string //优先级列表
}

func newcompareline(n string, c1 string, c2 string, l []string) *compareline {
	ret := new(compareline)
	ret.name = n
	ret.card1 = c1
	ret.card2 = c2
	ret.cardlist = l
	return ret
}

//卡牌优先级列表
type CardTable struct {
	Table     []compareline
	NameIndex map[string][]int
}

func NewCardTable() *CardTable {
	ret := new(CardTable)
	ret.Table = make([]compareline, 0, 3)
	ret.NameIndex = make(map[string][]int)
	return ret
}

//向优先级列表中插入一行
func (a *CardTable) AddList(line string) bool {
	name_content := strings.Split(line, ":")
	if len(name_content) < 2 {
		fmt.Println("the list is too short")
		return false
	}
	name := name_content[0]
	content := name_content[1]
	contents := strings.Split(content, " ")
	contents = tools.WipeOutBlank(contents)
	fmt.Println(contents)
	if len(contents) < 1 {
		fmt.Println("the list is too short")
		return false
	}
	mode := strings.Split(contents[0], "|")
	fmt.Println(mode)
	if len(mode) < 2 {
		fmt.Println("sytax error")
	}
	card1 := mode[0]
	card2 := mode[1]
	newline := *newcompareline(name, card1, card2, contents[1:])
	a.Table = append(a.Table, newline)
	index, exist := a.NameIndex[name]
	if exist == false {
		newindex := make([]int, 0, 3)
		newindex = append(newindex, len(a.Table)-1)
		a.NameIndex[name] = newindex
	} else {
		index = append(index, len(a.Table)-1)
		a.NameIndex[name] = index
	}
	return true
}

//判断两张牌是否能按照模式串进行匹配
func MatchMode(c1 cards.Card, c2 cards.Card, m1 string, m2 string) (bool, bool) {
	_, _, ok1 := cards.BreakRegEx(m1)
	_, _, ok2 := cards.BreakRegEx(m2)
	if ok1 == false || ok2 == false {
		fmt.Println("ERROR!!")
		return false, false
	}
	list := cards.NewCardList()
	list.AddCard(c1)
	list.AddCard(c2)
	regex := cards.NewRegEx("(" + m1 + ")(" + m2 + ")")
	code, ok := list.Match(*regex)
	if code == 0 {
		return true, ok
	} else {
		return false, ok
	}
}

//比较两张牌的大小
//1 表示 c1>c2 0 表示 c1==c2 -1表示 c1<c2 2表示无法比较
func (a *CardTable) Compare(c1 cards.Card, c2 cards.Card, standard string, table NameTable) int {
	index, exist := a.NameIndex[standard]
	if exist == false {
		fmt.Println("the standard ", standard, " doesn't exist")
		return 2
	}
	for i := 0; i < len(index); i++ {
		var list compareline = a.Table[index[i]]
		normal, ok := MatchMode(c1, c2, list.card1, list.card2)
		if normal == true && ok == true {
			var match1 int = len(list.cardlist)
			var match2 int = len(list.cardlist)
			for j := 0; j < len(list.cardlist); j++ {
				colour, point, suitable := cards.BreakRegEx(list.cardlist[j])
				colourexpression := strings.Split(colour, "%")
				pointexpression := strings.Split(point, "%")
				colour = ""
				point = ""
				for k := 0; k < len(colourexpression); k++ {
					if k%2 == 0 {
						colour = colour + colourexpression[k]
					} else {
						t := table.GetType(colourexpression[k])
						switch t {
						case "int":
							value := table.GetInt(colourexpression[k])
							colour = colour + cards.Int2Colour(value)
						case "string":
							colour = colour + table.GetString(colourexpression[k])
						default:
							fmt.Println("no variable")
						}
					}
				}
				for k := 0; k < len(pointexpression); k++ {
					if k%2 == 0 {
						point = point + pointexpression[k]
					} else {
						t := table.GetType(pointexpression[k])
						switch t {
						case "int":
							value := table.GetInt(pointexpression[k])
							point = point + cards.Int2Point(value)
						case "string":
							point = point + table.GetString(pointexpression[k])
						default:
							fmt.Println("no variable")
						}
					}
				}
				if suitable == true {
					if match1 >= len(list.cardlist) && c1.Match(colour, point) == true {
						fmt.Println(c1.ToString(), "->", colour, point)
						match1 = j
						if match2 < len(list.cardlist) {
							break
						}
					}
					if match2 >= len(list.cardlist) && c2.Match(colour, point) == true {
						fmt.Println(c2.ToString(), "->", colour, point)
						match2 = j
						if match1 < len(list.cardlist) {
							break
						}
					}
				}
			}
			switch {
			case match1 < match2:
				return 1
			case match1 > match2:
				return -1
			}
		}
	}
	return 0
}
