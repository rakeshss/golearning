package main

import (
	"fmt"
)

func main() {
	var i int
	i = 32
	fmt.Println(i)

	firstName := "Rakesh"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	real, imaginary := real(c), imag(c)
	fmt.Println(real, imaginary)
}
