package main

func main(){
	cards := deck {"one","two"}
	cards = append(cards,newCard())

	// print logic moved to deck.go
	cards.print()
}

func newCard() string{
	return "Five of Diamonds"
}

// to run this program: go run main.go deck.go 