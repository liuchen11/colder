package main

import (
	"codes/cards"
	"fmt"
	"os"
)

func TestStandardize(list cards.CardList) {
	list.Print()
	list.Standardize()
	list.Print()
}

func TestRandom() {
	list := cards.GenSetCards(1)
	list.Print()
	fmt.Println("---------------------")
	list.Disorganize()
	list.Print()
}

func main() {
	list := make([]cards.Card, 0, 10)
	list = append(list, *cards.NewCard("JB"), *cards.NewCard("HK"), *cards.NewCard("D0"), *cards.NewCard("CA"), *cards.NewCard("S5"))
	list = append(list, *cards.NewCardByIndex(1, 13), *cards.NewCardByIndex(0, 1))
	clist := cards.NewCardList()
	for i := 0; i < len(list); i++ {
		clist.AddCard(list[i])
	}
	clist.Print()
	table := clist.ToTable()
	for i := 0; i < 5; i++ {
		for j := 0; j < 14; j++ {
			fmt.Printf("%d ", table[i][j])
		}
		fmt.Printf("\n")
	}
	switch os.Args[1] {
	case "standardize":
		TestStandardize(*clist)
	case "random":
		TestRandom()
	default:
	}
}
