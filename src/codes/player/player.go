package player

/***************************************************************************************
**文件名：player.go
**包名称：player
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


import (
	"codes/cards"
	"fmt"
)

type Player struct {
	Holds cards.CardList
	Group int
}

func NewPlayer(group int) *Player {
	ret := new(Player)
	ret.Group = group
	ret.Holds = *cards.NewCardList()
	return ret
}

//获取最后摸到的牌
func (a *Player) Top_Card() cards.Card {
	return a.Holds.Cards[a.Holds.Length-1]
}

//寻找玩家手中的顺子等牌型
func (a *Player) Find(base string, length int, times int, strict bool) bool {
	start := cards.FindAllCards(base)
	table := a.Holds.ToTable()
	for i := 0; i < len(start); i++ {
		fmt.Println("cards->", start[i].ToString())
		var fit bool = true
		colour := start[i].GetColourIndex()
		point := start[i].GetPointIndex()
		if strict == true {
			for j, k := point, 0; k < length; k++ {
				if table[colour][j] < times {
					fit = false
					break
				}
				if j == 13 {
					j = 1
				} else {
					j++
				}
			}
		} else {
			var sum [14]int
			for i := 0; i < 14; i++ {
				sum[i] = 0
				for j := 1; j < 5; j++ {
					sum[i] += table[j][i]
				}
			}
			for j, k := point, 0; k < length; k++ {
				if sum[j] < times {
					fit = false
					break
				}
				if j == 13 {
					j = 1
				} else {
					j++
				}
			}
		}
		if fit == true {
			return true
		}
	}
	return false
}

//搜索一张牌，如果不存在，则返回-1
func Search(items []cards.Card, item cards.Card) int {
	for i := 0; i < len(items); i++ {
		if cards.Equals(items[i], item) == true {
			return i
		}
	}
	return -1
}

//传牌
func DeliverCards(sender *Player, receiver *Player, cardlist []cards.Card) bool {
	copy := make([]cards.Card, 0, 10)
	copy = append(copy, sender.Holds.Cards...)
	for i := 0; i < len(cardlist); i++ {
		index := Search(copy, cardlist[i])
		if index == -1 {
			return false
		}
		if index == len(copy)-1 {
			copy = copy[:index]
		} else {
			copy = append(copy[:index], copy[index+1:]...)
		}
	}
	sender.Holds.Cards = copy
	sender.Holds.Length = len(copy)
	receiver.Holds.Cards = append(receiver.Holds.Cards, cardlist...)
	receiver.Holds.Length = receiver.Holds.Length + len(cardlist)
	return true
}

