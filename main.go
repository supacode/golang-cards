package main

import "fmt"


func main(){
	cards := []string{
			newCard("Ace of Spade"),
			newCard("Five of Diamonds"),
		}

	for i, card := range cards {
		fmt.Println(i + 1, card)
	}
	
}

func newCard(card string) string {
	return card
}