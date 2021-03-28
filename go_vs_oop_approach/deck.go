package main

import "fmt"

type deck []string 

func (d deck) print() {
	// loop the slice
	for i,card := range d {
		fmt.Println(i, card)
	}
}