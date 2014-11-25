package parser

import (
	"bufio"
	"codes/cards"
	"codes/player"
	"codes/tools"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type TopModule struct {
	filename      string
	nametable     *NameTable
	functable     *FuncTable
	cardlisttable *CardListTable
	state         int
	pack_num      int
	player_num    int
	players       []player.Player
}

var TopLevel TopModule

func NewTopModule(name string) *TopModule {
	ret := new(TopModule)
	ret.filename = name
	ret.nametable = NewNameTable()
	ret.functable = new(FuncTable) //待修改
	ret.state = 0
	ret.pack_num = 0
	ret.player_num = 0
	ret.players = make([]player.Player, 0, 3)
	return ret
}

func (a *TopModule) GenParts() {
	result := SplitFile(a.filename)
	if result == false {
		fmt.Println("there are some problems spliting the .rule file")
		a.state = -1
	} else {
		a.state = 1
	}
}

func (a *TopModule) ParseConfig() {
	if a.state < 0 {
		fmt.Println("can't recover from the error!")
		return
	}
	fin, err := os.Open(a.filename + ".config")
	if err != nil {
		fmt.Println("there are some problems opening the config file")
		a.state = -2
	}
	reader := bufio.NewReader(fin)
	for {
		buffer, _, e := reader.ReadLine()
		if e == io.EOF {
			break
		} else if e != nil {
			fmt.Println("there are some problems loading the config file")
			a.state = -2
			return
		}
		buf := string(buffer)
		parts := strings.Split(buf, "=")
		if len(parts) == 2 {
			switch {
			case strings.Contains(buf, "cards"):
				number, er := strconv.ParseInt(parts[1], 10, 32)
				if er == nil {
					a.pack_num = int(number)
				}
			case strings.Contains(buf, "players"):
				number, er := strconv.ParseInt(parts[1], 10, 32)
				if er == nil {
					a.player_num = int(number)
				}
			default:
			}
		}
	}
	if a.pack_num <= 0 || a.player_num <= 0 {
		fmt.Println("there are some errors in your config file")
		a.state = -2
	} else {
		for i := 0; i <= a.player_num; i++ {
			toadd := *player.NewPlayer(i)
			a.players = append(a.players, toadd)
		}
	}
	return
}

func (a *TopModule) ParseVar() {
	if a.state < 0 {
		fmt.Println("can't recover from the error!")
		return
	}
	file := a.filename + ".var"
	p := NewParser(file)
	s, _ := GetBlockStmt(p.Content())
	flag := s.Exe(a.nametable, a.functable)
	switch flag {
	case false:
		fmt.Println("Some Errors in var block!")
		a.state = -3
	default:
		a.state = 3
	}
}

func (a *TopModule) ParseFunc() {
	if a.state < 0 {
		fmt.Println("can't recover from the error!")
		return
	}
	file := a.filename + "func"
	p := NewParser(file)
	var flag bool
	a.functable, flag = GetFuncTable(p.Content())
	switch flag {
	case true:
		a.state = 4
	default:
		fmt.Println("there are some problems in the Func part")
		a.state = -4
	}
}

func (a *TopModule) ParseCard() {
	if a.state < 0 {
		fmt.Println("can't recover from the error!")
		return
	}
	file := a.filename + ".card"
	fin, err := os.Open(file)
	if err != nil {
		fmt.Println("can't open the file in the Card part")
		a.state = -5
		return
	}
	defer fin.Close()
	reader := bufio.NewReader(fin)
	for {
		buffer, _, e := reader.ReadLine()
		if e == io.EOF {
			break
		} else if e != nil {
			fmt.Println("there are some problems while loading the file in the Card part")
			a.state = -5
			return
		}
		buf := string(buffer)
		if !tools.IsBlank(buf) {
			result := a.cardlisttable.AddList(buf)
			if result == false {
				fmt.Println("there is syntax error in the Card part")
			}
		}
	}
	a.state = 5
}

func (a *TopModule) ParseMode() {
	if a.state < 0 {
		fmt.Println("can't recover from the error!")
		return
	}
	file := a.filename + ".mode"
	fin, err := os.Open(file)
	if err != nil {
		fmt.Println("can't open the file in the Mode part")
		a.state = -6
		return
	}
	defer fin.Close()
	reader := bufio.NewReader(fin)
	for {
		buffer, _, e := reader.ReadLine()
		if e == io.EOF {
			break
		} else if e != nil {
			fmt.Println("there are some problems while loading the file in the Mode part")
			a.state = -6
			return
		}
		buf := string(buffer)
		if !tools.IsBlank(buf) {
			result := a.cardlisttable.AddMode(buf)
			if result == false {
				fmt.Println("there is syntax error in the Card part")
			}
		}
	}
	a.state = 6
}

func (a *TopModule) ParseBody() {
	file := a.filename + ".body"
	p := NewParser(file)
	s, _ := GetBlockStmt(p.Content())
	flag := s.Exe(a.nametable, a.functable)
	switch flag {
	case true:
		a.state = 7
	default:
		fmt.Println("there are some problems in the Body part")
		a.state = -7
	}
}

func (a *TopModule) ParseEnd() {
	file := a.filename + ".end"
	p := NewParser(file)
	s, _ := GetBlockStmt(p.Content())
	flag := s.Exe(a.nametable, a.functable)
	switch flag {
	case true:
		a.state = 8
	default:
		fmt.Println("there are some problems in the End part")
		a.state = -8
	}
}

func (a *TopModule) Get_Cards(num int) string {
	if num > a.player_num {
		return ""
	}
	var play_cards string = ""
	for i := 0; i < a.players[num].Holds.Length; i++ {
		play_cards = play_cards + a.players[num].Holds.Cards[i].ToString()
	}
	return play_cards
}

func (a *TopModule) Get_CardSum(num int) int {
	if num > a.player_num {
		return 0
	} else {
		return a.players[num].Holds.Length
	}
}

func (a *TopModule) Top_Card(num int) string {
	if num > a.player_num {
		return ""
	} else {
		number := a.Get_CardSum(num)
		if number > 0 {
			return a.players[num].Holds.Cards[number-1].ToString()
		} else {
			return ""
		}
	}
}

func (a *TopModule) Find(num int, base string, length int, times int, strict bool) bool {
	if num > a.player_num {
		return false
	} else {
		return a.players[num].Find(base, length, times, strict)
	}
}

func (a *TopModule) Match(cardlist string, expression string) bool {
	regex := cards.NewRegEx(expression)
	list := cards.NewCardList()
	for i := 0; 2*i+1 < len(cardlist); i++ {
		toadd := *cards.NewCard(cardlist[2*i : 2*i+2])
		list.AddCard(toadd)
	}
	code, result := list.Match(*regex)
	if code == 0 && result == true {
		return true
	} else {
		return false
	}
}

func (a *TopModule) Compare(card1 string, card2 string) int {
	list1 := cards.NewCardList()
	list2 := cards.NewCardList()
	for i := 0; 2*i+1 < len(card1); i++ {
		toadd := *cards.NewCard(card1[2*i : 2*i+2])
		list1.AddCard(toadd)
	}
	for i := 0; 2*i+1 < len(card2); i++ {
		toadd := *cards.NewCard(card2[2*i : 2*i+2])
		list2.AddCard(toadd)
	}
	codes, ok := a.cardlisttable.Compare(*list1, *list2, *a.nametable)
	if ok == false {
		fmt.Println("some error while comparing ", card1, " and ", card2)
		return 2
	} else {
		return codes
	}
}

func (a *TopModule) Deliver(sender int, receiver int, card string, show bool) bool {
	if show == false { ///
		return false ///
	} ///
	list := make([]cards.Card, 0, 5)
	for i := 0; 2*i+1 < len(card); i++ {
		toadd := *cards.NewCard(card[2*i : 2*i+2])
		list = append(list, toadd)
	}
	if sender > a.player_num || receiver > a.player_num {
		return false
	} else {
		return player.DeliverCards(&a.players[sender], &a.players[receiver], list)
	}
}

func (a *TopModule) Next_Round() {
	//TODO
}

func (a *TopModule) play_card(num int, tocompare string) int {
	for {
		var post string = "JB" //处理出牌
		if post == "" {
			return -1
		}
		posted := *cards.NewCardList()
		for i := 0; 2*i+1 < len(post); i++ {
			toadd := *cards.NewCard(post[2*i : 2*i+2])
			posted.AddCard(toadd)
		}
		list := *cards.NewCardList()
		for i := 0; 2*i+1 < len(tocompare); i++ {
			toadd := *cards.NewCard(tocompare[2*i : 2*i+2])
			list.AddCard(toadd)
		}
		result, ok := a.cardlisttable.Post(posted, list, *a.nametable)
		if result == 1 && ok == true {
			return 1
		}
	}
}

func (a *TopModule) Enable(num int) {
	//TODO
}

func (a *TopModule) Disable(num int) {
	//TODO
}

func (a *TopModule) State_Point(num int, set []int, candrop bool) int {
	//TODO
	return 0
}

func (a *TopModule) Show_Card(num int, length int) string {
	//TODO
	return ""
}

func (a *TopModule) Show_Public(message string) {
	//TODO
}

func (a *TopModule) Show_Private(num int, message string) {
	//TODO
}

func (a *TopModule) Add_Score(num int, score int) {
	//TODO
}

func (a *TopModule) Begin() {
	for i := 0; i < len(a.players); i++ {
		a.players[i].Holds = *cards.NewCardList()
	}
	a.players[0].Holds = *cards.GenSetCards(a.pack_num)
	a.players[0].Holds.Disorganize()
}
