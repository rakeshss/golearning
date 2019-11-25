package main

import ( "fmt")

func main(){
    vertices := make(map[string]int)
	
	vertices["Triangle"] = 3
	vertices["Square"] = 4
	vertices["Pentagon"] = 5
	
	fmt.Println(vertices)
	fmt.Println(vertices["Triangle"])
	
	delete(vertices,"Square")
	fmt.Println(vertices)
}
