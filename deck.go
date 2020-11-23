package main

import "fmt"

type deck []string

func (d deck) show() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}