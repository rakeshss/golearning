package main

import "fmt"

func main() {
	a()
	b()
}

func a() {
	i := 0
	defer fmt.Println(i)
	i = 100
	return
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}
