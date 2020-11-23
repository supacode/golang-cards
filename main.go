package main


func main(){

	deck := createDeck()

	hand, deckRemainder := deal(deck, 5) 


	hand.show()
	deckRemainder.show()	

}