package main

import (
	"fmt"
)

func main() {
	// implicitly constant
	const c = 3
	fmt.Println(c)

	fmt.Println(c + 1)
	fmt.Println(c + 1.3)

	// explicitly constant
	const cc int = 3
	fmt.Println(cc)

	fmt.Println(cc + 1)
	//fmt.Println(cc + 1.3)  -- > this will give error
}
