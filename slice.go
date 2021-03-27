package main

import("fmt")

func main(){

	cards := []string {"one","two"}
	cards = append(cards,newCard())
	fmt.Println(cards)

	// loop the slice
	for i,card := range(cards){
		fmt.Println(i, card)
	}
}

func newCard() string{
	return "Five of Diamonds"
}