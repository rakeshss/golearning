package main

import("fmt")

func main(){

	cards := []string {"one","two"}
	cards = append(cards,newCard())
	fmt.Println(cards)
}

func newCard() string{
	return "Five of Diamonds"
}