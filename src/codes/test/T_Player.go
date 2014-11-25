package main

import (
	"codes/cards"
	"codes/player"
	"fmt"
	"os"
)

func T_find() {
	list := make([]cards.Card, 0, 10)
	list = append(list, *cards.NewCard("HA"), *cards.NewCard("HK"), *cards.NewCard("D0"), *cards.NewCard("CA"), *cards.NewCard("S5"))
	list = append(list, *cards.NewCardByIndex(1, 13), *cards.NewCardByIndex(0, 1))
	clist := cards.NewCardList()
	for i := 0; i < len(list); i++ {
		clist.AddCard(list[i])
	}
	clist.Print()
	player := *player.NewPlayer(0)
	player.Holds = *clist
	result1 := player.Find(".K", 3, 1, true)
	fmt.Println(result1)
}

func T_Deliver() {
	list1 := make([]cards.Card, 0, 10)
	list1 = append(list1, *cards.NewCard("HA"), *cards.NewCard("HK"), *cards.NewCard("D0"), *cards.NewCard("CA"), *cards.NewCard("S5"))
	list1 = append(list1, *cards.NewCardByIndex(1, 13), *cards.NewCardByIndex(0, 1))
	clist1 := cards.NewCardList()
	for i := 0; i < len(list1); i++ {
		clist1.AddCard(list1[i])
	}
	clist1.Print()
	player1 := *player.NewPlayer(0)
	player1.Holds = *clist1

	list2 := make([]cards.Card, 0, 10)
	list2 = append(list2, *cards.NewCard("HA"), *cards.NewCard("HK"), *cards.NewCard("D0"), *cards.NewCard("CA"), *cards.NewCard("S5"))
	list2 = append(list2, *cards.NewCardByIndex(1, 13), *cards.NewCardByIndex(0, 1))
	clist2 := cards.NewCardList()
	for i := 0; i < len(list2); i++ {
		clist2.AddCard(list2[i])
	}
	clist2.Print()
	player2 := *player.NewPlayer(0)
	player2.Holds = *clist2

	list3 := make([]cards.Card, 0, 10)
	list3 = append(list3, *cards.NewCard("HA"), *cards.NewCard("HK"), *cards.NewCard("D0"), *cards.NewCard("CA"), *cards.NewCard("S5"))
	list3 = append(list3, *cards.NewCardByIndex(1, 13), *cards.NewCardByIndex(0, 1))
	player1.Holds.Print()
	fmt.Println("-------")
	player2.Holds.Print()
	player.DeliverCards(&player1, &player2, list3)
	fmt.Println("END")
	player1.Holds.Print()
	fmt.Println("-------")
	player2.Holds.Print()
}

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "find":
			T_find()
		case "deliver":
			T_Deliver()
		default:
		}
	}
}
