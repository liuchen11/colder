package cards

/***************************************************************************************
**文件名：CardList.go
**包名称：cards
**创建日期：2013-12-21
**作者：The Colder
**版本：v1.0
**支持平台：windows/Linux/Mac OSX
**说明：一个支持自己编写规则集的在线卡牌游戏，2013年软件工程大作业
***************************************************************************************/


/*******************************************************************
*文件名：...所在包...
*创建日期....
*作者:the colder
*版本...
*支持平台 Windows/Mac OSX/Linux
*说明
********************************************************************/

import (
	"fmt"
	"math/rand"
	"time"
)

//牌组类型
type CardList struct {
	Cards  []Card
	Length int
}

//新建牌组
func NewCardList() *CardList {
	ret := new(CardList)
	ret.Cards = make([]Card, 0, 5)
	ret.Length = 0
	return ret
}

//向牌组中加上一张牌
func (a *CardList) AddCard(toadd Card) {
	a.Length++
	a.Cards = append(a.Cards, toadd)
}

//两个牌组进行合并
func (a *CardList) Merge(list CardList) {
	a.Cards = append(a.Cards, list.Cards...)
	a.Length = a.Length + list.Length
}

//打印牌组
func (a *CardList) Print() {
	for i := 0; i < a.Length; i++ {
		a.Cards[i].Print()
		fmt.Print(" ")
	}
	fmt.Print("\n")
}

//标准化所用的辅助类表示某一点数的牌的数量和花色分布
type list struct {
	sum int
	num [4]int
}

//添加某一类型的牌的时候，list表的变化
func (a *list) add(index int) {
	a.sum++
	a.num[index]++
}

/*
按照正则表达式的标准将牌组中牌的顺序标准化
顺序为 先大王小王 然后按照张数进行排序 张数相同按照A-K进行排序 同一点数按照 花色进行排序
采用哈希表和桶排序的方法
*/
func (a *CardList) Standardize() {
	jb := 0
	js := 0
	max := 0
	boxes := make([]list, 13, 13) //构建13个桶表示13个整数
	for i := 0; i < len(a.Cards); i++ {
		c := a.Cards[i].GetColourIndex()
		p := a.Cards[i].GetPointIndex()
		if c == 0 { //统计王牌
			if p == 2 {
				jb++
			} else {
				js++
			}
		} else { //一般牌
			boxes[p-1].add(c - 1)
			if boxes[p-1].sum > max {
				max = boxes[p-1].sum
			}
		}
	}
	a.Cards = make([]Card, 0, 10)
	for i := 0; i < jb; i++ {
		a.Cards = append(a.Cards, *NewCardByIndex(0, 2))
	}
	for i := 0; i < js; i++ {
		a.Cards = append(a.Cards, *NewCardByIndex(0, 1))
	}
	for i := max; i > 0; i-- {
		for j := 0; j < 13; j++ {
			if boxes[j].sum == i {
				for k := 0; k < 4; k++ {
					for l := 0; l < boxes[j].num[k]; l++ {
						a.Cards = append(a.Cards, *NewCardByIndex(k+1, j+1))
					}
				}
			}
		}
	}
}

//牌组的正则表达式匹配
func (a *CardList) Match(regex RegEx) (int, bool) {
	//将牌组转化成标准形式
	a.Standardize()
	//消除{}内是确定数字的情况,将正则表达式转化为标准形式
	return regex.Match(*a)
}

//生成几副牌乱牌
func GenSetCards(number int) *CardList {
	list := NewCardList()
	for i := 0; i < number; i++ {
		var toadd Card
		toadd = *NewCardByIndex(0, 2)
		list.AddCard(toadd)
		toadd = *NewCardByIndex(0, 1)
		list.AddCard(toadd)
		for j := 1; j <= 4; j++ {
			for k := 1; k <= 13; k++ {
				toadd = *NewCardByIndex(j, k)
				list.AddCard(toadd)
			}
		}
	}
	return list
}

//从哈希表中查找对应的命中桶，不命中，则返回-1
func find(start int, hashmap []int) int {
	for i := start; ; i++ {
		i = i % len(hashmap)
		if hashmap[i] == 0 {
			return i
		}
		if i == (start+len(hashmap)-1)%len(hashmap) {
			return -1
		}
	}
}

//打乱排序
func (a *CardList) Disorganize() {
	var total int = a.Length
	var pre []Card = a.Cards
	a.Cards = make([]Card, 0, 10)
	hashmap := make([]int, total, total)
	for i := 0; i < total; i++ {
		hashmap[i] = 0
	}
	for i := 0; i < total; i++ {
		present := time.Now().Nanosecond()
		rand.Seed(int64(present) + int64(i*i))
		start := rand.Int() % total
		hit := find(start, hashmap)
		a.Cards = append(a.Cards, pre[hit])
		hashmap[hit] = 1
	}
}

//转化为Hash表
func (a *CardList) ToTable() [5][14]int {
	var table [5][14]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 14; j++ {
			table[i][j] = 0
		}
	}
	for i := 0; i < len(a.Cards); i++ {
		colour := a.Cards[i].GetColourIndex()
		point := a.Cards[i].GetPointIndex()
		table[colour][point]++
	}
	return table
}

//寻找一张牌，如果不存在，则返回-1
func (a *CardList) Search(item Card) int {
	for i := 0; i < len(a.Cards); i++ {
		if Equals(a.Cards[i], item) == true {
			return i
		}
	}
	return -1
}

