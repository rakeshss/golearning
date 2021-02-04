package main

import (
	"fmt"
)

func main() {
	var firstName *string = new(string)
	fmt.Println(firstName)

	*firstName = "Rakesh"
	fmt.Println(*firstName)
}
