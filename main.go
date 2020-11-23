package main

func main(){
	cards := deck {
		newCard("Ace of Spade"),
		newCard("King of Hearts"),
		newCard("Five of Diamonds"),
	}

	cards = append(cards, newCard("Joker"))


	cards.show()

}


func newCard(cardName string) string {
	return cardName
}