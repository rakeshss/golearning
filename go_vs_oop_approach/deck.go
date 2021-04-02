package main

import "fmt"

type deck []string 

// this is called reciever Go/method/function receiver
// d here is similar to this, self we use in python/java script.

func (d deck) print() {
	// loop the slice
	for i,card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {

	cards := deck{}

	cardsSuits := []string{"Spades","Diamonds","Hearts"}
	cardsValues := []string{"one","two","three","four:"}

	for _,suits := range cardsSuits {
		for _,values := range cardsValues{
			cards = append(cards,suits + " of " + values)
		}
	}

	return cards
}