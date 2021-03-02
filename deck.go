package main

import "fmt"

type deck []string

func createDeck() deck {

	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			newCard := value + " of " + suite
			cards = append(cards, newCard)
		}
	}

	return cards

}

func (d deck) show() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}
