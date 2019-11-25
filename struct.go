package main

import (
	"fmt"
)

type employee struct {
	name string
	empId int 
}

func main(){
	p := employee{name: "Rakesh", empId : 100}
	fmt.Println(p);
	fmt.Println(p.name);
	fmt.Println(p.empId);
}
